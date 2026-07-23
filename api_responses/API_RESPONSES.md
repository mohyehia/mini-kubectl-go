# API Responses

## Version Command

### Request
```shell
curl --cacert ca.crt --cert client.crt --key client.key https://127.0.0.1:6443/version
```

### Response
```json
{
  "major": "1",
  "minor": "35",
  "emulationMajor": "1",
  "emulationMinor": "35",
  "minCompatibilityMajor": "1",
  "minCompatibilityMinor": "34",
  "gitVersion": "v1.35.5+k3s1",
  "gitCommit": "6a4781ad53ee5cad273bedcd9462ae36ac97d798",
  "gitTreeState": "clean",
  "buildDate": "2026-05-20T11:15:10Z",
  "goVersion": "go1.25.9",
  "compiler": "gc",
  "platform": "linux/amd64"
}
```
---
## APIs Command
### Request
```shell
curl --cacert ca.crt --cert client.crt --key client.key https://127.0.0.1:6443/api
```

### Response
```json
{
  "kind": "APIGroupList",
  "apiVersion": "v1",
  "groups": [
    {
      "name": "apiregistration.k8s.io",
      "versions": [
        {
          "groupVersion": "apiregistration.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "apiregistration.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "apps",
      "versions": [
        {
          "groupVersion": "apps/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "apps/v1",
        "version": "v1"
      }
    },
    {
      "name": "events.k8s.io",
      "versions": [
        {
          "groupVersion": "events.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "events.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "authentication.k8s.io",
      "versions": [
        {
          "groupVersion": "authentication.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "authentication.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "authorization.k8s.io",
      "versions": [
        {
          "groupVersion": "authorization.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "authorization.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "autoscaling",
      "versions": [
        {
          "groupVersion": "autoscaling/v2",
          "version": "v2"
        },
        {
          "groupVersion": "autoscaling/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "autoscaling/v2",
        "version": "v2"
      }
    },
    {
      "name": "batch",
      "versions": [
        {
          "groupVersion": "batch/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "batch/v1",
        "version": "v1"
      }
    },
    {
      "name": "certificates.k8s.io",
      "versions": [
        {
          "groupVersion": "certificates.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "certificates.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "networking.k8s.io",
      "versions": [
        {
          "groupVersion": "networking.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "networking.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "policy",
      "versions": [
        {
          "groupVersion": "policy/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "policy/v1",
        "version": "v1"
      }
    },
    {
      "name": "rbac.authorization.k8s.io",
      "versions": [
        {
          "groupVersion": "rbac.authorization.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "rbac.authorization.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "storage.k8s.io",
      "versions": [
        {
          "groupVersion": "storage.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "storage.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "admissionregistration.k8s.io",
      "versions": [
        {
          "groupVersion": "admissionregistration.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "admissionregistration.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "apiextensions.k8s.io",
      "versions": [
        {
          "groupVersion": "apiextensions.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "apiextensions.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "scheduling.k8s.io",
      "versions": [
        {
          "groupVersion": "scheduling.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "scheduling.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "coordination.k8s.io",
      "versions": [
        {
          "groupVersion": "coordination.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "coordination.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "node.k8s.io",
      "versions": [
        {
          "groupVersion": "node.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "node.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "discovery.k8s.io",
      "versions": [
        {
          "groupVersion": "discovery.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "discovery.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "resource.k8s.io",
      "versions": [
        {
          "groupVersion": "resource.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "resource.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "flowcontrol.apiserver.k8s.io",
      "versions": [
        {
          "groupVersion": "flowcontrol.apiserver.k8s.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "flowcontrol.apiserver.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "gateway.networking.k8s.io",
      "versions": [
        {
          "groupVersion": "gateway.networking.k8s.io/v1",
          "version": "v1"
        },
        {
          "groupVersion": "gateway.networking.k8s.io/v1beta1",
          "version": "v1beta1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "gateway.networking.k8s.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "helm.cattle.io",
      "versions": [
        {
          "groupVersion": "helm.cattle.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "helm.cattle.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "k3s.cattle.io",
      "versions": [
        {
          "groupVersion": "k3s.cattle.io/v1",
          "version": "v1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "k3s.cattle.io/v1",
        "version": "v1"
      }
    },
    {
      "name": "hub.traefik.io",
      "versions": [
        {
          "groupVersion": "hub.traefik.io/v1alpha1",
          "version": "v1alpha1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "hub.traefik.io/v1alpha1",
        "version": "v1alpha1"
      }
    },
    {
      "name": "traefik.io",
      "versions": [
        {
          "groupVersion": "traefik.io/v1alpha1",
          "version": "v1alpha1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "traefik.io/v1alpha1",
        "version": "v1alpha1"
      }
    },
    {
      "name": "metrics.k8s.io",
      "versions": [
        {
          "groupVersion": "metrics.k8s.io/v1beta1",
          "version": "v1beta1"
        }
      ],
      "preferredVersion": {
        "groupVersion": "metrics.k8s.io/v1beta1",
        "version": "v1beta1"
      }
    }
  ]
}
```
