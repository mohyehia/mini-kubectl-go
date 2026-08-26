package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/mohyehia/mini-kubectl/pkg/k8s/resources"
)

func PrintDeployments(out io.Writer, responseBody []byte, namespace, outputFormat string) error {
	var resourceList resources.ResourceList[resources.Deployment]
	if err := json.Unmarshal(responseBody, &resourceList); err != nil {
		return fmt.Errorf("failed to unmarshal resource list: %w", err)
	}
	if hasNoResources(namespace, len(resourceList.Items)) {
		return nil
	}

	switch outputFormat {
	case "default", "table":
		return printDeploymentsDefaultFormat(out, &resourceList)
	case "json":
		return printJsonFormat(out, responseBody)
	case "yaml", "yml":
		return printYamlFormat(out, responseBody)
	default:
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}
}

func printDeploymentsDefaultFormat(out io.Writer, resourceList *resources.ResourceList[resources.Deployment]) error {
	// minWidth: 0, tabWidth: 8, padding: 2 (spaces between columns), padChar: ' ', flags: 0
	writer := tabwriter.NewWriter(out, 0, 8, 2, ' ', 0)
	// 1. Print Column Headers (separated by tabs)
	_, err := fmt.Fprintln(writer, "NAME\tREADY\tUP-TO-DATE\tAVAILABLE\tAGE")
	if err != nil {
		return err
	}

	// 2. Print Each Deployment Row
	for _, deploy := range resourceList.Items {
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
