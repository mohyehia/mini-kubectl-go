# PODS API

### Request
```shell
curl --cacert ca.crt --cert client.crt --key client.key https://127.0.0.1:26443/api/v1/namespaces/default/pods
```

### Response
```json
{
  "kind": "PodList",
  "apiVersion": "v1",
  "metadata": {
    "resourceVersion": "4834"
  },
  "items": [
    {
      "metadata": {
        "name": "nginx",
        "namespace": "default",
        "uid": "8f1152d3-2c0c-499a-8c18-9be44c6b0a84",
        "resourceVersion": "4791",
        "generation": 1,
        "creationTimestamp": "2026-08-08T15:16:02Z",
        "labels": {
          "run": "nginx"
        },
        "managedFields": [
          {
            "manager": "kubectl-run",
            "operation": "Update",
            "apiVersion": "v1",
            "time": "2026-08-08T15:16:01Z",
            "fieldsType": "FieldsV1",
            "fieldsV1": {
              "f:metadata": {
                "f:labels": {
                  ".": {},
                  "f:run": {}
                }
              },
              "f:spec": {
                "f:containers": {
                  "k:{\"name\":\"nginx\"}": {
                    ".": {},
                    "f:image": {},
                    "f:imagePullPolicy": {},
                    "f:name": {},
                    "f:resources": {},
                    "f:terminationMessagePath": {},
                    "f:terminationMessagePolicy": {}
                  }
                },
                "f:dnsPolicy": {},
                "f:enableServiceLinks": {},
                "f:restartPolicy": {},
                "f:schedulerName": {},
                "f:securityContext": {},
                "f:terminationGracePeriodSeconds": {}
              }
            }
          },
          {
            "manager": "k3s",
            "operation": "Update",
            "apiVersion": "v1",
            "time": "2026-08-08T15:16:30Z",
            "fieldsType": "FieldsV1",
            "fieldsV1": {
              "f:status": {
                "f:conditions": {
                  "k:{\"type\":\"ContainersReady\"}": {
                    ".": {},
                    "f:lastProbeTime": {},
                    "f:lastTransitionTime": {},
                    "f:observedGeneration": {},
                    "f:status": {},
                    "f:type": {}
                  },
                  "k:{\"type\":\"Initialized\"}": {
                    ".": {},
                    "f:lastProbeTime": {},
                    "f:lastTransitionTime": {},
                    "f:observedGeneration": {},
                    "f:status": {},
                    "f:type": {}
                  },
                  "k:{\"type\":\"PodReadyToStartContainers\"}": {
                    ".": {},
                    "f:lastProbeTime": {},
                    "f:lastTransitionTime": {},
                    "f:observedGeneration": {},
                    "f:status": {},
                    "f:type": {}
                  },
                  "k:{\"type\":\"PodScheduled\"}": {
                    "f:observedGeneration": {}
                  },
                  "k:{\"type\":\"Ready\"}": {
                    ".": {},
                    "f:lastProbeTime": {},
                    "f:lastTransitionTime": {},
                    "f:observedGeneration": {},
                    "f:status": {},
                    "f:type": {}
                  }
                },
                "f:containerStatuses": {},
                "f:hostIP": {},
                "f:hostIPs": {},
                "f:observedGeneration": {},
                "f:phase": {},
                "f:podIP": {},
                "f:podIPs": {
                  ".": {},
                  "k:{\"ip\":\"192.168.194.16\"}": {
                    ".": {},
                    "f:ip": {}
                  }
                },
                "f:startTime": {}
              }
            },
            "subresource": "status"
          }
        ]
      },
      "spec": {
        "volumes": [
          {
            "name": "kube-api-access-dq9sh",
            "projected": {
              "sources": [
                {
                  "serviceAccountToken": {
                    "expirationSeconds": 3607,
                    "path": "token"
                  }
                },
                {
                  "configMap": {
                    "name": "kube-root-ca.crt",
                    "items": [
                      {
                        "key": "ca.crt",
                        "path": "ca.crt"
                      }
                    ]
                  }
                },
                {
                  "downwardAPI": {
                    "items": [
                      {
                        "path": "namespace",
                        "fieldRef": {
                          "apiVersion": "v1",
                          "fieldPath": "metadata.namespace"
                        }
                      }
                    ]
                  }
                }
              ],
              "defaultMode": 420
            }
          }
        ],
        "containers": [
          {
            "name": "nginx",
            "image": "nginx",
            "resources": {},
            "volumeMounts": [
              {
                "name": "kube-api-access-dq9sh",
                "readOnly": true,
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount"
              }
            ],
            "terminationMessagePath": "/dev/termination-log",
            "terminationMessagePolicy": "File",
            "imagePullPolicy": "Always"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "serviceAccountName": "default",
        "serviceAccount": "default",
        "nodeName": "orbstack",
        "securityContext": {},
        "schedulerName": "default-scheduler",
        "tolerations": [
          {
            "key": "node.kubernetes.io/not-ready",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          },
          {
            "key": "node.kubernetes.io/unreachable",
            "operator": "Exists",
            "effect": "NoExecute",
            "tolerationSeconds": 300
          }
        ],
        "priority": 0,
        "enableServiceLinks": true,
        "preemptionPolicy": "PreemptLowerPriority"
      },
      "status": {
        "observedGeneration": 1,
        "phase": "Running",
        "conditions": [
          {
            "type": "PodReadyToStartContainers",
            "observedGeneration": 1,
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2026-08-08T15:16:30Z"
          },
          {
            "type": "Initialized",
            "observedGeneration": 1,
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2026-08-08T15:16:02Z"
          },
          {
            "type": "Ready",
            "observedGeneration": 1,
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2026-08-08T15:16:30Z"
          },
          {
            "type": "ContainersReady",
            "observedGeneration": 1,
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2026-08-08T15:16:30Z"
          },
          {
            "type": "PodScheduled",
            "observedGeneration": 1,
            "status": "True",
            "lastProbeTime": null,
            "lastTransitionTime": "2026-08-08T15:16:02Z"
          }
        ],
        "hostIP": "192.168.139.2",
        "hostIPs": [
          {
            "ip": "192.168.139.2"
          },
          {
            "ip": "fd07:b51a:cc66::2"
          }
        ],
        "podIP": "192.168.194.16",
        "podIPs": [
          {
            "ip": "192.168.194.16"
          }
        ],
        "startTime": "2026-08-08T15:16:02Z",
        "containerStatuses": [
          {
            "name": "nginx",
            "state": {
              "running": {
                "startedAt": "2026-08-08T15:16:29Z"
              }
            },
            "lastState": {},
            "ready": true,
            "restartCount": 0,
            "image": "nginx:latest",
            "imageID": "docker-pullable://nginx@sha256:8541484afbc9c8a5a8a99b379568ebbc957f658583ec9448fc43104229c03cf8",
            "containerID": "docker://c96b0b854df2b42e65100166c7aa45e0d92e16a55025c7be87126c186ad37dbe",
            "started": true,
            "resources": {},
            "volumeMounts": [
              {
                "name": "kube-api-access-dq9sh",
                "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                "readOnly": true,
                "recursiveReadOnly": "Disabled"
              }
            ]
          }
        ],
        "qosClass": "BestEffort"
      }
    }
  ]
} 
```