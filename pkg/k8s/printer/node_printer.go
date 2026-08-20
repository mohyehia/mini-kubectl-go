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

func PrintNodes(out io.Writer, responseBody []byte, outputFormat string) error {
	var resourceList resources.ResourceList[resources.Node]
	if err := json.Unmarshal(responseBody, &resourceList); err != nil {
		return fmt.Errorf("failed to unmarshal resource list: %w", err)
	}
	switch outputFormat {
	case "default", "table":
		return printNodesDefaultFormat(out, &resourceList)
	case "json":
		return printNodesJsonFormat(out, responseBody)
	case "yaml", "yml":
		return printNodesYamlFormat(out, responseBody)
	default:
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}
}

func printNodesDefaultFormat(out io.Writer, nodes *resources.ResourceList[resources.Node]) error {
	// minWidth: 0, tabWidth: 8, padding: 2 (spaces between columns), padChar: ' ', flags: 0
	tabWriter := tabwriter.NewWriter(out, 0, 8, 2, ' ', 0)
	// 1. Print Column Headers (separated by tabs)
	_, err := fmt.Fprintln(tabWriter, "NAME\tSTATUS\tROLES\tAGE\tVERSION")
	if err != nil {
		return err
	}

	// 2. Print Each Node Row
	for _, node := range nodes.Items {
		status := getNodeStatus(node)
		role := getNodeRole(node)
		_, err := fmt.Fprintf(tabWriter, "%s\t%s\t%s\t%s\t%s\n",
			node.Metadata.Name,
			status,
			role,
			node.Metadata.GetAge(),
			node.Status.Info.KubeletVersion,
		)
		if err != nil {
			return err
		}
	}
	return tabWriter.Flush()
}

func printNodesJsonFormat(out io.Writer, responseBody []byte) error {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, responseBody, "", "  "); err != nil {
		return fmt.Errorf("failed to format JSON: %w", err)
	}
	_, err := out.Write(prettyJSON.Bytes())
	return err
}

func printNodesYamlFormat(out io.Writer, responseBody []byte) error {
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

func getNodeStatus(node resources.Node) string {
	status := "Unknown"
	conditions := node.Status.Conditions
	for _, condition := range conditions {
		if condition.Type == "Ready" {
			if condition.Status == "True" {
				return "Ready"
			}
			return "NotReady"
		}
	}
	return status
}

func getNodeRole(node resources.Node) string {
	role := "<none>"
	labels := node.Metadata.Labels
	if _, exists := labels["node-role.kubernetes.io/control-plane"]; exists {
		role = "control-plane"
	} else if _, exists := labels["node-role.kubernetes.io/master"]; exists {
		role = "control-plane"
	}
	return role
}
