# Deployments API

### Request
```shell
curl --cacert ca.crt --cert client.crt --key client.key https://127.0.0.1:26443/apis/apps/v1/namespaces/default/deployments
```

### Response
```json
{
  "kind": "DeploymentList",
  "apiVersion": "apps/v1",
  "metadata": {
    "resourceVersion": "10290"
  },
  "items": [
    {
      "metadata": {
        "name": "nginx-deployment",
        "namespace": "default",
        "uid": "8ed8a225-b538-4d27-adc2-fa6c9bf6baeb",
        "resourceVersion": "10227",
        "generation": 1,
        "creationTimestamp": "2026-08-15T10:04:57Z",
        "labels": {
          "app": "nginx"
        },
        "annotations": {
          "deployment.kubernetes.io/revision": "1",
          "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{},\"labels\":{\"app\":\"nginx\"},\"name\":\"nginx-deployment\",\"namespace\":\"default\"},\"spec\":{\"replicas\":2,\"selector\":{\"matchLabels\":{\"app\":\"nginx\"}},\"template\":{\"metadata\":{\"labels\":{\"app\":\"nginx\"}},\"spec\":{\"containers\":[{\"image\":\"nginx:latest\",\"name\":\"nginx\",\"ports\":[{\"containerPort\":80}],\"resources\":{\"limits\":{\"cpu\":\"500m\",\"memory\":\"512Mi\"},\"requests\":{\"cpu\":\"250m\",\"memory\":\"256Mi\"}}}]}}}}\n"
        },
        "managedFields": [
          {
            "manager": "kubectl-client-side-apply",
            "operation": "Update",
            "apiVersion": "apps/v1",
            "time": "2026-08-15T10:04:57Z",
            "fieldsType": "FieldsV1",
            "fieldsV1": {
              "f:metadata": {
                "f:annotations": {
                  ".": {},
                  "f:kubectl.kubernetes.io/last-applied-configuration": {}
                },
                "f:labels": {
                  ".": {},
                  "f:app": {}
                }
              },
              "f:spec": {
                "f:progressDeadlineSeconds": {},
                "f:replicas": {},
                "f:revisionHistoryLimit": {},
                "f:selector": {},
                "f:strategy": {
                  "f:rollingUpdate": {
                    ".": {},
                    "f:maxSurge": {},
                    "f:maxUnavailable": {}
                  },
                  "f:type": {}
                },
                "f:template": {
                  "f:metadata": {
                    "f:labels": {
                      ".": {},
                      "f:app": {}
                    }
                  },
                  "f:spec": {
                    "f:containers": {
                      "k:{\"name\":\"nginx\"}": {
                        ".": {},
                        "f:image": {},
                        "f:imagePullPolicy": {},
                        "f:name": {},
                        "f:ports": {
                          ".": {},
                          "k:{\"containerPort\":80,\"protocol\":\"TCP\"}": {
                            ".": {},
                            "f:containerPort": {},
                            "f:protocol": {}
                          }
                        },
                        "f:resources": {
                          ".": {},
                          "f:limits": {
                            ".": {},
                            "f:cpu": {},
                            "f:memory": {}
                          },
                          "f:requests": {
                            ".": {},
                            "f:cpu": {},
                            "f:memory": {}
                          }
                        },
                        "f:terminationMessagePath": {},
                        "f:terminationMessagePolicy": {}
                      }
                    },
                    "f:dnsPolicy": {},
                    "f:restartPolicy": {},
                    "f:schedulerName": {},
                    "f:securityContext": {},
                    "f:terminationGracePeriodSeconds": {}
                  }
                }
              }
            }
          },
          {
            "manager": "k3s",
            "operation": "Update",
            "apiVersion": "apps/v1",
            "time": "2026-08-15T10:05:01Z",
            "fieldsType": "FieldsV1",
            "fieldsV1": {
              "f:metadata": {
                "f:annotations": {
                  "f:deployment.kubernetes.io/revision": {}
                }
              },
              "f:status": {
                "f:availableReplicas": {},
                "f:conditions": {
                  ".": {},
                  "k:{\"type\":\"Available\"}": {
                    ".": {},
                    "f:lastTransitionTime": {},
                    "f:lastUpdateTime": {},
                    "f:message": {},
                    "f:reason": {},
                    "f:status": {},
                    "f:type": {}
                  },
                  "k:{\"type\":\"Progressing\"}": {
                    ".": {},
                    "f:lastTransitionTime": {},
                    "f:lastUpdateTime": {},
                    "f:message": {},
                    "f:reason": {},
                    "f:status": {},
                    "f:type": {}
                  }
                },
                "f:observedGeneration": {},
                "f:readyReplicas": {},
                "f:replicas": {},
                "f:terminatingReplicas": {},
                "f:updatedReplicas": {}
              }
            },
            "subresource": "status"
          }
        ]
      },
      "spec": {
        "replicas": 2,
        "selector": {
          "matchLabels": {
            "app": "nginx"
          }
        },
        "template": {
          "metadata": {
            "labels": {
              "app": "nginx"
            }
          },
          "spec": {
            "containers": [
              {
                "name": "nginx",
                "image": "nginx:latest",
                "ports": [
                  {
                    "containerPort": 80,
                    "protocol": "TCP"
                  }
                ],
                "resources": {
                  "limits": {
                    "cpu": "500m",
                    "memory": "512Mi"
                  },
                  "requests": {
                    "cpu": "250m",
                    "memory": "256Mi"
                  }
                },
                "terminationMessagePath": "/dev/termination-log",
                "terminationMessagePolicy": "File",
                "imagePullPolicy": "Always"
              }
            ],
            "restartPolicy": "Always",
            "terminationGracePeriodSeconds": 30,
            "dnsPolicy": "ClusterFirst",
            "securityContext": {},
            "schedulerName": "default-scheduler"
          }
        },
        "strategy": {
          "type": "RollingUpdate",
          "rollingUpdate": {
            "maxUnavailable": "25%",
            "maxSurge": "25%"
          }
        },
        "revisionHistoryLimit": 10,
        "progressDeadlineSeconds": 600
      },
      "status": {
        "observedGeneration": 1,
        "replicas": 2,
        "updatedReplicas": 2,
        "readyReplicas": 2,
        "availableReplicas": 2,
        "terminatingReplicas": 0,
        "conditions": [
          {
            "type": "Available",
            "status": "True",
            "lastUpdateTime": "2026-08-15T10:05:01Z",
            "lastTransitionTime": "2026-08-15T10:05:01Z",
            "reason": "MinimumReplicasAvailable",
            "message": "Deployment has minimum availability."
          },
          {
            "type": "Progressing",
            "status": "True",
            "lastUpdateTime": "2026-08-15T10:05:01Z",
            "lastTransitionTime": "2026-08-15T10:04:57Z",
            "reason": "NewReplicaSetAvailable",
            "message": "ReplicaSet \"nginx-deployment-6f888c676c\" has successfully progressed."
          }
        ]
      }
    }
  ]
}
```