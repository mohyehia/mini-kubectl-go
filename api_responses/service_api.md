# Services List API

### Request
```shell
curl --cacert ca.crt --cert client.crt --key client.key https://127.0.0.1:26443/api/v1/namespaces/default/services
```

### Response
```json
{
  "kind": "ServiceList",
  "apiVersion": "v1",
  "metadata": {
    "resourceVersion": "4623"
  },
  "items": [
    {
      "metadata": {
        "name": "kubernetes",
        "namespace": "default",
        "uid": "069e0090-75d5-4251-bf3f-6f0133accdf3",
        "resourceVersion": "66",
        "creationTimestamp": "2026-07-23T16:35:04Z",
        "labels": {
          "component": "apiserver",
          "provider": "kubernetes"
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
                  "f:component": {},
                  "f:provider": {}
                }
              },
              "f:spec": {
                "f:clusterIP": {},
                "f:internalTrafficPolicy": {},
                "f:ipFamilyPolicy": {},
                "f:ports": {
                  ".": {},
                  "k:{\"port\":443,\"protocol\":\"TCP\"}": {
                    ".": {},
                    "f:name": {},
                    "f:port": {},
                    "f:protocol": {},
                    "f:targetPort": {}
                  }
                },
                "f:sessionAffinity": {},
                "f:type": {}
              }
            }
          }
        ]
      },
      "spec": {
        "ports": [
          {
            "name": "https",
            "protocol": "TCP",
            "port": 443,
            "targetPort": 26443
          }
        ],
        "clusterIP": "192.168.194.129",
        "clusterIPs": [
          "192.168.194.129"
        ],
        "type": "ClusterIP",
        "sessionAffinity": "None",
        "ipFamilies": [
          "IPv4"
        ],
        "ipFamilyPolicy": "SingleStack",
        "internalTrafficPolicy": "Cluster"
      },
      "status": {
        "loadBalancer": {}
      }
    }
  ]
}
```