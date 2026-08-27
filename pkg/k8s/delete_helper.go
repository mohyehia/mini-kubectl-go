package k8s

import (
	"context"
	"fmt"
	"io"
	"time"
)

func WaitForDeletion(ctx context.Context, out io.Writer, c *Client, kind ResourceKind, namespace, resourceName string, pollInterval time.Duration) error {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context canceled or timed out while waiting for resource deletion: %w", ctx.Err())
		case <-ticker.C:
			exists, err := c.Exists(ctx, kind, namespace, resourceName)
			if err != nil {
				return fmt.Errorf("error checking resource existence: %w", err)
			}
			if !exists {
				// 404 Not Found, resource has been deleted successfully!
				return nil
			}
		}
	}
}
