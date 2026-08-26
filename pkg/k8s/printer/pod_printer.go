package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/mohyehia/mini-kubectl/pkg/k8s/resources"
)

func PrintPods(out io.Writer, responseBody []byte, namespace, outputFormat string) error {
	var resourceList resources.ResourceList[resources.Pod]
	if err := json.Unmarshal(responseBody, &resourceList); err != nil {
		return fmt.Errorf("failed to unmarshal resource list: %w", err)
	}
	if hasNoResources(namespace, len(resourceList.Items)) {
		return nil
	}

	switch outputFormat {
	case "default", "table":
		return printPodsDefaultFormat(out, &resourceList)
	case "json":
		return printJsonFormat(out, responseBody)
	case "yaml", "yml":
		return printYamlFormat(out, responseBody)
	default:
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}
}

func printPodsDefaultFormat(out io.Writer, resourceList *resources.ResourceList[resources.Pod]) error {
	// minWidth: 0, tabWidth: 8, padding: 2 (spaces between columns), padChar: ' ', flags: 0
	tabWriter := tabwriter.NewWriter(out, 0, 8, 2, ' ', 0)
	// 1. Print Column Headers (separated by tabs)
	_, err := fmt.Fprintln(tabWriter, "NAME\tREADY\tSTATUS\tRESTARTS\tAGE")
	if err != nil {
		return err
	}

	// 2. Print Each Pod Row
	for _, pod := range resourceList.Items {
		readyCount := 0
		restartCount := 0
		totalCount := len(pod.Status.ContainerStatuses)
		for _, status := range pod.Status.ContainerStatuses {
			if status.Ready {
				readyCount++
			}
			restartCount += int(status.RestartCount)
		}
		ready := fmt.Sprintf("%d/%d", readyCount, totalCount)
		_, err := fmt.Fprintf(tabWriter, "%s\t%s\t%s\t%d\t%s\n",
			pod.Metadata.Name,
			ready,
			pod.Status.Phase,
			restartCount,
			pod.Metadata.GetAge(),
		)
		if err != nil {
			return err
		}
	}
	return tabWriter.Flush()
}
