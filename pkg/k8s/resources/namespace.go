package resources

type Namespace struct {
	Metadata ObjectMetadata  `json:"metadata"`
	Status   NamespaceStatus `json:"status"`
}

type NamespaceStatus struct {
	Phase string `json:"phase"`
}
