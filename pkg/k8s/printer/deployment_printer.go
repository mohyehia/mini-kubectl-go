package printer

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/mohyehia/mini-kubectl/pkg/k8s/resources"
)

func PrintDeployments(out io.Writer, deployments *resources.ResourceList[resources.Deployment]) error {
	// minWidth: 0, tabWidth: 8, padding: 2 (spaces between columns), padChar: ' ', flags: 0
	writer := tabwriter.NewWriter(out, 0, 8, 2, ' ', 0)
	// 1. Print Column Headers (separated by tabs)
	_, err := fmt.Fprintln(writer, "NAME\tREADY\tUP-TO-DATE\tAVAILABLE\tAGE")
	if err != nil {
		return err
	}

	// 2. Print Each Deployment Row
	for _, deploy := range deployments.Items {
		ready := fmt.Sprintf("%d/%d", deploy.Status.ReadyReplicas, deploy.Status.Replicas)
		_, err := fmt.Fprintf(writer, "%s\t%s\t%d\t%d\t%s\n",
			deploy.Metadata.Name,
			ready,
			deploy.Status.UpdatedReplicas,
			deploy.Status.AvailableReplicas,
			deploy.Metadata.GetAge(),
		)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
