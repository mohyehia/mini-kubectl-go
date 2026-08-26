package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/mohyehia/mini-kubectl/pkg/k8s/resources"
)

func PrintNamespaces(out io.Writer, responseBody []byte, outputFormat string) error {
	var namespaces resources.ResourceList[resources.Namespace]
	if err := json.Unmarshal(responseBody, &namespaces); err != nil {
		return fmt.Errorf("failed to unmarshal resource list: %w", err)
	}
	switch outputFormat {
	case "default", "table":
		return printNamespacesDefaultFormat(out, &namespaces)
	case "json":
		return printJsonFormat(out, responseBody)
	case "yaml", "yml":
		return printYamlFormat(out, responseBody)
	default:
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}
}

func printNamespacesDefaultFormat(out io.Writer, namespaces *resources.ResourceList[resources.Namespace]) error {
	// minWidth: 0, tabWidth: 8, padding: 2 (spaces between columns), padChar: ' ', flags: 0
	tabWriter := tabwriter.NewWriter(out, 0, 8, 2, ' ', 0)
	// 1. Print Column Headers (separated by tabs)
	_, err := fmt.Fprintln(tabWriter, "NAME\tSTATUS\tAGE")
	if err != nil {
		return err
	}

	// 2. Print Each Namespace Row
	for _, ns := range namespaces.Items {
		status := ns.Status.Phase
		if status == "" {
			status = "<unknown>"
		}
		_, err := fmt.Fprintf(tabWriter, "%s\t%s\t%s\n",
			ns.Metadata.Name,
			status,
			ns.Metadata.GetAge(),
		)
		if err != nil {
			return err
		}
	}
	return tabWriter.Flush()
}
