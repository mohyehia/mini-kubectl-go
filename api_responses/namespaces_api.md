# Namespaces API

### Request
```shell
curl --cacert ca.crt --cert client.crt --key client.key https://127.0.0.1:26443/api/v1/namespaces
```

### Response
```json
{
  "kind": "NamespaceList",
  "apiVersion": "v1",
  "metadata": {
    "resourceVersion": "4485"
  },
  "items": [
    {
      "metadata": {
        "name": "default",
        "uid": "4fc1d442-046a-4abb-a614-9307cf76963b",
        "resourceVersion": "18",
        "creationTimestamp": "2026-07-23T16:35:04Z",
        "labels": {
          "kubernetes.io/metadata.name": "default"
        },
        "managedFields": [
          {
            "manager": "k3s",
            "operation": "Update",
            "apiVersion": "v1",
            "time": "2026-07-23T16:35:04Z",
            "fieldsType": "FieldsV1",
            "fieldsV1": {
              "f:metadata": {
                "f:labels": {
                  ".": {},
                  "f:kubernetes.io/metadata.name": {}
                }
              }
            }
          }
        ]
      },
      "spec": {
        "finalizers": [
          "kubernetes"
        ]
      },
      "status": {
        "phase": "Active"
      }
    },
    {
      "metadata": {
        "name": "kube-node-lease",
        "uid": "38f96c6a-9fb0-4f1d-a270-51af6cee75ff",
        "resourceVersion": "25",
        "creationTimestamp": "2026-07-23T16:35:04Z",
        "labels": {
          "kubernetes.io/metadata.name": "kube-node-lease"
        },
        "managedFields": [
          {
            "manager": "k3s",
            "operation": "Update",
            "apiVersion": "v1",
            "time": "2026-07-23T16:35:04Z",
            "fieldsType": "FieldsV1",
            "fieldsV1": {
              "f:metadata": {
                "f:labels": {
                  ".": {},
                  "f:kubernetes.io/metadata.name": {}
                }
              }
            }
          }
        ]
      },
      "spec": {
        "finalizers": [
          "kubernetes"
        ]
      },
      "status": {
        "phase": "Active"
      }
    },
    {
      "metadata": {
        "name": "kube-public",
        "uid": "0eb67e95-c59b-480e-80a1-fb886dcaf8da",
        "resourceVersion": "9",
        "creationTimestamp": "2026-07-23T16:35:04Z",
        "labels": {
          "kubernetes.io/metadata.name": "kube-public"
        },
        "managedFields": [
          {
            "manager": "k3s",
            "operation": "Update",
            "apiVersion": "v1",
            "time": "2026-07-23T16:35:04Z",
            "fieldsType": "FieldsV1",
            "fieldsV1": {
              "f:metadata": {
                "f:labels": {
                  ".": {},
                  "f:kubernetes.io/metadata.name": {}
                }
              }
            }
          }
        ]
      },
      "spec": {
        "finalizers": [
          "kubernetes"
        ]
      },
      "status": {
        "phase": "Active"
      }
    },
    {
      "metadata": {
        "name": "kube-system",
        "uid": "c045e392-f2c8-4ea3-b218-94d37bd6c5b6",
        "resourceVersion": "5",
        "creationTimestamp": "2026-07-23T16:35:04Z",
        "labels": {
          "kubernetes.io/metadata.name": "kube-system"
        },
        "managedFields": [
          {
            "manager": "k3s",
            "operation": "Update",
            "apiVersion": "v1",
            "time": "2026-07-23T16:35:04Z",
            "fieldsType": "FieldsV1",
            "fieldsV1": {
              "f:metadata": {
                "f:labels": {
                  ".": {},
                  "f:kubernetes.io/metadata.name": {}
                }
              }
            }
          }
        ]
      },
      "spec": {
        "finalizers": [
          "kubernetes"
        ]
      },
      "status": {
        "phase": "Active"
      }
    }
  ]
} 
```