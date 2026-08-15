package resources

type Deployment struct {
	Metadata ObjectMetadata   `json:"metadata"`
	Status   DeploymentStatus `json:"status"`
}

type DeploymentStatus struct {
	Replicas          int32 `json:"replicas"`
	UpdatedReplicas   int32 `json:"updatedReplicas"`
	ReadyReplicas     int32 `json:"readyReplicas"`
	AvailableReplicas int32 `json:"availableReplicas"`
}
