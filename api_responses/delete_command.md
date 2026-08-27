To get the **maximum systems engineering learning value** out of `delete`, we shouldn't just fire off a simple fire-and-forget `DELETE` request.

Instead, we will implement **Synchronous Wait & Polling Semantics** (similar to `kubectl delete --wait=true`).

---

## 🧠 Why This Yields Maximum Learning

When you issue a `DELETE` request to Kubernetes, the API server doesn't immediately delete the resource. It sets a `deletionTimestamp` and hands control over to background controllers (e.g., Kubelet terminating container processes gracefully).

By implementing a wait loop, you will learn:

1. **HTTP `DELETE` Semantics:** Sending deletion calls and interpreting `200 OK` vs. `202 Accepted` vs. `Status` objects.
2. **Polling & Context Cancellation:** Using Go's `time.Ticker` alongside `select` and `context.WithTimeout` to handle timeouts gracefully.
3. **HTTP `404 Not Found` as a Success Signal:** Understanding state convergence in eventual consistency models.

---

## 🏗️ High-Level Sequence

```text
 User Execution: stridectl delete pod my-pod -n default
                         │
                         ▼
             1. Issue HTTP DELETE Request
                         │
        ┌────────────────┴────────────────┐
        ▼                                 ▼
   200 / 202 OK                     404 Not Found
 (Marked for Deletion)            (Resource Doesn't Exist)
        │                                 │
        ▼                                 ▼
2. Start Polling Loop              Return Error / Exit
   (Ticker every 500ms)
        │
        ├──> GET /api/v1/namespaces/default/pods/my-pod
        │       │
        │       ├──> Returns 200 OK  ──> Keep Polling
        │       └──> Returns 404     ──> Success! Exit Loop
        │
        └──> Timeout Exceeded (e.g., 30s) ──> Context Cancelled Return Error

```

---

## 🛠️ Step 1: Add `DeleteRaw` and `Exists` Methods to `client.go`

Add these methods to your mTLS HTTP client in `pkg/k8s/client.go`:

```go
package k8s

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// DeleteRaw sends an HTTP DELETE request to the given Kubernetes API URL.
func (c *Client) DeleteRaw(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create DELETE request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("DELETE request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("resource not found")
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("DELETE returned unexpected status %d: %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// Exists checks if a resource still exists by issuing a GET request.
// Returns false if the server responds with 404 Not Found.
func (c *Client) Exists(ctx context.Context, url string) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return false, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, fmt.Errorf("unexpected status check code: %d", resp.StatusCode)
}

```

---

## 🛠️ Step 2: Implement `WaitForDeletion` Polling Engine

Add this helper inside `pkg/k8s/delete.go`. This logic uses Go's `select` statement to balance a ticker, a timeout context, and context cancellation:

```go
package k8s

import (
	"context"
	"fmt"
	"io"
	"time"
)

// WaitForDeletion polls the API server endpoint until the resource returns 404 (deleted) or context times out.
func WaitForDeletion(ctx context.Context, out io.Writer, client *Client, url string, pollInterval time.Duration) error {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timed out waiting for resource deletion: %w", ctx.Err())

		case <-ticker.C:
			exists, err := client.Exists(ctx, url)
			if err != nil {
				return fmt.Errorf("error polling resource status: %w", err)
			}

			if !exists {
				// 404 Not Found received -> Deletion successfully confirmed!
				return nil
			}
		}
	}
}

```

---

## 🛠️ Step 3: Build `cmd/delete.go`

Now wire up Cobra with a `--wait` and `--timeout` flag:

```go
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"your-project/pkg/k8s"
)

var (
	waitFlag    bool
	timeoutFlag time.Duration
)

var deleteCmd = &cobra.Command{
	Use:   "delete TYPE NAME",
	Short: "Delete a resource by type and name",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		resourceType := args[0]
		resourceName := args[1]

		// 1. Resolve resource kind and construct target URL
		info, err := k8s.Registry.Resolve(resourceType)
		if err != nil {
			return err
		}

		targetURL := k8s.BuildSingleResourceURL(info, namespaceFlag, resourceName)

		// 2. Setup Context with Timeout
		ctx, cancel := context.WithTimeout(context.Background(), timeoutFlag)
		defer cancel()

		// 3. Issue the HTTP DELETE
		fmt.Printf("deleting %s %q...\n", info.Kind, resourceName)
		_, err = client.DeleteRaw(ctx, targetURL)
		if err != nil {
			return err
		}

		// 4. Handle Wait Flag logic
		if waitFlag {
			fmt.Printf("waiting up to %s for %s %q to terminate...\n", timeoutFlag, info.Kind, resourceName)
			if err := k8s.WaitForDeletion(ctx, os.Stdout, client, targetURL, 500*time.Millisecond); err != nil {
				return err
			}
		}

		fmt.Printf("%s %q deleted\n", info.Kind, resourceName)
		return nil
	},
}

func init() {
	deleteCmd.Flags().BoolVar(&waitFlag, "wait", true, "Wait for resource to be fully deleted from cluster")
	deleteCmd.Flags().DurationVar(&timeoutFlag, "timeout", 30*time.Second, "The length of time to wait before giving up on deletion")
	rootCmd.AddCommand(deleteCmd)
}

```

---

## 🖥️ Output in Terminal

Running `stridectl delete pod nginx-pod -n default --wait`:

```text
deleting POD "nginx-pod"...
waiting up to 30s for POD "nginx-pod" to terminate...
POD "nginx-pod" deleted

```
