package printer

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/mohyehia/mini-kubectl/pkg/k8s/resources"
)

func PrintPods(out io.Writer, pods *resources.ResourceList[resources.Pod]) error {

	// minWidth: 0, tabWidth: 8, padding: 2 (spaces between columns), padChar: ' ', flags: 0
	tabWriter := tabwriter.NewWriter(out, 0, 8, 2, ' ', 0)
	// 1. Print Column Headers (separated by tabs)
	_, err := fmt.Fprintln(tabWriter, "NAME\tREADY\tSTATUS\tRESTARTS\tAGE")
	if err != nil {
		return err
	}

	// 2. Print Each Pod Row
	for _, pod := range pods.Items {
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
