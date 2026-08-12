package resources

import (
	"fmt"
	"time"
)

func (om *ObjectMetadata) GetAge() string {
	if om.CreationTimestamp.IsZero() {
		return "<unknown>"
	}
	duration := time.Since(om.CreationTimestamp)
	if duration.Hours() > 24 {
		return fmt.Sprintf("%dd", int(duration.Hours()/24))
	}
	if duration.Hours() >= 1 {
		return fmt.Sprintf("%dh", int(duration.Hours()))
	}
	if duration.Minutes() >= 1 {
		return fmt.Sprintf("%dm", int(duration.Minutes()))
	}
	return fmt.Sprintf("%ds", int(duration.Seconds()))
}
