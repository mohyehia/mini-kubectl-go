package printer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/mohyehia/mini-kubectl/pkg/k8s/resources"
	"gopkg.in/yaml.v3"
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
		return printNamespacesJsonFormat(out, responseBody)
	case "yaml", "yml":
		return printNamespacesYamlFormat(out, responseBody)
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

func printNamespacesJsonFormat(out io.Writer, responseBody []byte) error {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, responseBody, "", "  "); err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	_, err := out.Write(prettyJSON.Bytes())
	return err
}

func printNamespacesYamlFormat(out io.Writer, responseBody []byte) error {
	var yamlData any
	if err := json.Unmarshal(responseBody, &yamlData); err != nil {
		return fmt.Errorf("failed to parse JSON for YAML conversion: %w", err)
	}
	yamlBytes, err := yaml.Marshal(yamlData)
	if err != nil {
		return fmt.Errorf("failed to convert JSON to YAML: %w", err)
	}
	_, err = out.Write(yamlBytes)
	return err
}
