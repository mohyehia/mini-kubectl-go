package printer

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"github.com/mohyehia/mini-kubectl/pkg/k8s/resources"
)

func PrintServices(out io.Writer, responseBody []byte, namespace, outputFormat string) error {
	var resourceList resources.ResourceList[resources.Service]
	if err := json.Unmarshal(responseBody, &resourceList); err != nil {
		return fmt.Errorf("failed to unmarshal resource list: %w", err)
	}
	if hasNoResources(namespace, len(resourceList.Items)) {
		return nil
	}

	switch outputFormat {
	case "default", "table":
		return printServicesDefaultFormat(out, &resourceList)
	case "json":
		return printJsonFormat(out, responseBody)
	case "yaml", "yml":
		return printYamlFormat(out, responseBody)
	default:
		return fmt.Errorf("unsupported output format: %s", outputFormat)
	}
}

func printServicesDefaultFormat(out io.Writer, resourceList *resources.ResourceList[resources.Service]) error {
	// minWidth: 0, tabWidth: 8, padding: 2 (spaces between columns), padChar: ' ', flags: 0
	writer := tabwriter.NewWriter(out, 0, 8, 2, ' ', 0)
	// 1. Print Column Headers (separated by tabs)
	_, err := fmt.Fprintln(writer, "NAME\tTYPE\tCLUSTER-IP\tEXTERNAL-IP\tPORTS\tAGE")
	if err != nil {
		return err
	}

	// 2. Print Each Service Row
	for _, svc := range resourceList.Items {
		externalIP := getExternalIP(svc)

		portStrings := make([]string, 0, len(svc.Spec.Ports))
		for _, port := range svc.Spec.Ports {
			portStrings = append(portStrings, fmt.Sprintf("%d:%d/%s", port.Port, port.TargetPort, port.Protocol))
		}
		ports := strings.Join(portStrings, ",")

		_, err := fmt.Fprintf(writer, "%s\t%s\t%s\t%s\t%s\t%s\n",
			svc.Metadata.Name,
			svc.Spec.Type,
			svc.Spec.ClusterIP,
			externalIP,
			ports,
			svc.Metadata.GetAge(),
		)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func getExternalIP(svc resources.Service) string {
	if len(svc.Spec.ExternalIPs) > 0 {
		return strings.Join(svc.Spec.ExternalIPs, ",")
	}

	if svc.Spec.Type == "ExternalName" && svc.Spec.ExternalName != "" {
		return svc.Spec.ExternalName
	}

	if len(svc.Status.LoadBalancer.Ingress) > 0 {
		values := make([]string, 0, len(svc.Status.LoadBalancer.Ingress))
		for _, ingress := range svc.Status.LoadBalancer.Ingress {
			if ingress.IP != "" {
				values = append(values, ingress.IP)
			} else if ingress.Hostname != "" {
				values = append(values, ingress.Hostname)
			}
		}
		if len(values) > 0 {
			return strings.Join(values, ",")
		}
	}

	return "<none>"
}
