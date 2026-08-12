package resources

type ResourceList[T any] struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Items      []T    `json:"items"`
}
