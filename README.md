# mini-kubectl-go
Go application for mimicking some of kubectl commands just for learning fun 

## Commands for calling k8s api_server directly:

#### 1. Extract and decode the credentials into temp files
```shell
grep "certificate-authority-data" config | awk '{print $2}' | base64 -d > ca.crt
grep "client-certificate-data" config | awk '{print $2}' | base64 -d > client.crt
grep "client-key-data" config | awk '{print $2}' | base64 -d > client.key

```
#### 2. Make the authenticated call
```shell
curl --cacert ca.crt --cert client.crt --key client.key https://127.0.0.1:6443/version
```
### 3. Clean up the temp files
```shell
rm ca.crt client.crt client.key
```

## HLD
┌────────────────────────────────────────────────────────┐
│                      main.go (CLI)                     │
│  1. Cobra intercepts the --config flag path string.    │
│  2. Passes that string path down to pkg/k8s.           │
└───────────────────────────┬────────────────────────────┘
│
│ "Hey pkg/k8s, read this file path
│  and give me back an active client!"
▼
┌────────────────────────────────────────────────────────┐
│                    pkg/k8s (Engine)                    │
│  3. Reads the YAML file.                               │
│  4. Finds the active context, clusters, and users.     │
│  5. Spins up the authenticated mTLS HTTP connection.   │
│  6. Returns a ready-to-use Client Interface object.    │
└───────────────────────────┬────────────────────────────┘
│
│ "Here is your working Client.
│  You can use it for any command."
▼
┌────────────────────────────────────────────────────────┐
│                    Back to Cobra CLI                   │
│  7. Passes this single Client object directly to the   │
│     version, get, or delete command handlers.          │
└────────────────────────────────────────────────────────┘


## 📋 The 4-Step Extraction Sequence
[Start]
│
▼
1. Read File ──────> Reads raw bytes from disk (e.g., 'config.yaml')
   │
   ▼
2. Parse YAML ────> Ingests bytes via yaml.v3 into the KubeConfig struct
   │
   ▼
3. Find Context ──> Locates the object matching 'current-context' string
   │
   ▼
4. Extract Keys ──> Pulls Target Server URL, CA Data, Cert Data, & Key Data
   │
   ▼
   [Next Phase: Network Setup]
