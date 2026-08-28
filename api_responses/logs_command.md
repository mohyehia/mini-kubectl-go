# Logs Command Specification

### 📐 Architecture & Data Flow

```text
  [Terminal / SIGINT]
          │
          ▼
┌─────────────────────────────────────────────────────────────────────────┐
│ signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)      │
└────────────────────────────────────┬────────────────────────────────────┘
                                     │ Passes Cancellable Context
                                     ▼
┌─────────────────────────────────────────────────────────────────────────┐
│ http.NewRequestWithContext(ctx, "GET", "/api/v1/.../log?follow=true")   │
└────────────────────────────────────┬────────────────────────────────────┘
                                     │ Initiates HTTP Transport
                                     ▼
┌─────────────────────────────────────────────────────────────────────────┐
│ Kubernetes API Server (Transfer-Encoding: chunked stream)               │
└────────────────────────────────────┬────────────────────────────────────┘
                                     │ Continuous Byte Stream
                                     ▼
┌─────────────────────────────────────────────────────────────────────────┐
│ io.Copy(os.Stdout, resp.Body) ───> Real-time Log Output to Terminal     │
└─────────────────────────────────────────────────────────────────────────┘

```

---

### 🌐 HTTP Endpoint Specification

| Attribute              | Value                                                                                                                  |
|------------------------|------------------------------------------------------------------------------------------------------------------------|
| **Path**               | `/api/v1/namespaces/{namespace}/pods/{podName}/log`                                                                    |
| **HTTP Method**        | `GET`                                                                                                                  |
| **Query Parameters**   | `follow=true` (enables continuous streaming), `container=<name>` (optional), `tailLines=100` (optional initial buffer) |
| **Transport Encoding** | `Transfer-Encoding: chunked` (plain text payload, no JSON wrapping)                                                    |

---

### 🧠 Core Component Specifications

**1. Context-Bound Network Layer**

* Use `http.NewRequestWithContext(ctx, ...)` to bind the HTTP connection's lifetime directly to a Go `context.Context`.
* When the context is active, `c.httpClient.Do(req)` executes the request and returns an open `resp.Body` reader without waiting for an EOF.

**2. Zero-Allocation Streaming Pipeline**

* Pass `resp.Body` directly into `io.Copy(os.Stdout, resp.Body)`.
* This bypasses intermediate array allocations and JSON unmarshaling, piping raw chunked bytes from the TCP socket buffer straight to standard output as soon as frames arrive.

**3. Graceful Signal Handling & Socket Cleanup (AC 6.2)**

* Intercept `os.Interrupt` (Ctrl+C) and `syscall.SIGTERM` using `signal.NotifyContext`.
* When a signal is caught:
1. The signal interceptor calls the context's internal `cancel()` function.
2. `net/http` detects `ctx.Done()` and immediately triggers a socket shutdown (`TCP FIN/RST`) to the API server.
3. `resp.Body.Read()` unblocks and returns `net.ErrClosed` / `context.Canceled`.
4. `io.Copy` exits instantly, unblocking the main execution thread without leaving zombie goroutines running in the background.



---

### 🛠️ Execution & State Handling Logic

1. **Parameter Validation:** Verify that `podName` is supplied and resolve the target namespace via `namespaceFlag`.
2. **Stream Initialization:** Execute the HTTP request. If the server returns a non-200 status (e.g., `404 Not Found` or `400 Container Not Specified`), consume the initial error body and exit immediately.
3. **Stream Copy:** Execute `io.Copy(os.Stdout, resp.Body)`.
4. **Exit Routing:**
* If `io.Copy` returns `nil` or `context.Canceled` because of `Ctrl+C`, treat it as a clean shutdown and return `nil` to Cobra.
* If `io.Copy` returns a socket read error while the context is still active, return the wrapped network error.