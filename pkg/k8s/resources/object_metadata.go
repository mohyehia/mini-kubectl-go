package resources

import "time"

type ObjectMetadata struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	CreationTimestamp time.Time         `json:"creationTimestamp"`
	Version           string            `json:"resourceVersion"`
	Labels            map[string]string `json:"labels"`
}
