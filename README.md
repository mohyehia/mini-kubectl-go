# mini-kubectl-go

A small Go CLI that mimics a subset of `kubectl` by calling the Kubernetes API directly for learning purposes.

## What works

- `version` — prints client and server version
- `get` — lists resources in `default`, `table`, `json`, or `yaml` output
- `delete` — deletes a resource by type and name, with wait and timeout support
- `describe` — currently a work in progress

## Supported resources

- Pods: `po`, `pod`, `pods`
- Services: `svc`, `service`, `services`
- Deployments: `deploy`, `deployment`, `deployments`
- Namespaces: `ns`, `namespace`, `namespaces`
- Nodes: `no`, `node`, `nodes`

## How it connects

- Reads kubeconfig from `~/.kube/config` by default
- Supports `-c, --config` to use a custom kubeconfig path
- Supports `-n, --namespace` to target a namespace
- Uses certificate-based mTLS credentials from kubeconfig

## Useful flags

- `get <resource> -o default|table|json|yaml`
- `delete <type> <name> --wait --timeout 30s`
- `describe (TYPE [NAME] | TYPE/NAME)` is available but still incomplete

## Quick examples

```shell
go run . version
go run . get pods
go run . get deployments -n kube-system -o yaml
go run . get svc -o json
go run . delete pod nginx --wait --timeout 30s
```

## Install / Build

```shell
go build -o mini-kubectl .
./mini-kubectl version
```

## Current limitations

- `describe` is still incomplete, and authentication currently supports kubeconfig entries with client certificate/key data only.

## Notes

- This project talks to the Kubernetes API directly instead of using `kubectl` or `client-go`

