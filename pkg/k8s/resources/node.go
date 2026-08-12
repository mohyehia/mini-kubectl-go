package resources

type Node struct {
	Metadata ObjectMetadata `json:"metadata"`
	Status   NodeStatus     `json:"status"`
}

type NodeStatus struct {
	Conditions []Condition `json:"conditions"`
	Info       NodeInfo    `json:"nodeInfo"`
}

type NodeInfo struct {
	OsImage         string `json:"osImage"`
	KubeletVersion  string `json:"kubeletVersion"`
	OperatingSystem string `json:"operatingSystem"`
}

type Condition struct {
	Type   string `json:"type"`
	Status string `json:"status"`
}
