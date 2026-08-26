package cmd

import (
	"fmt"
	"os"

	"github.com/mohyehia/mini-kubectl/pkg/k8s"
	"github.com/mohyehia/mini-kubectl/pkg/k8s/printer"
	"github.com/spf13/cobra"
)

var outputFormat string
var getCommand = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources",
	Long:  "Display one or many resources",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		if appState.k8sClient == nil {
			fmt.Println("Server info is not available at the moment.")
			return nil
		}

		err := validateOutputFormat(outputFormat)
		if err != nil {
			return err
		}

		resourceType := args[0]

		info, err := validateResourceType(resourceType)
		if err != nil {
			return err
		}
		responseBody, err := appState.k8sClient.GetResourceList(info.Kind, appState.namespace)
		if err != nil {
			return err
		}

		switch info.Kind {
		case k8s.NODE:
			return printer.PrintNodes(os.Stdout, responseBody, outputFormat)
		case k8s.NAMESPACE:
			return printer.PrintNamespaces(os.Stdout, responseBody, outputFormat)
		case k8s.POD:
			return printer.PrintPods(os.Stdout, responseBody, appState.namespace, outputFormat)
		case k8s.SERVICE:
			return printer.PrintServices(os.Stdout, responseBody, appState.namespace, outputFormat)
		case k8s.DEPLOYMENT:
			return printer.PrintDeployments(os.Stdout, responseBody, appState.namespace, outputFormat)
		default:
			return fmt.Errorf("unsupported resource type: %s", resourceType)
		}
	},
}

func init() {
	getCommand.Flags().StringVarP(&outputFormat, "output", "o", "default", "Output format. One of: json|yaml|table")
	rootCmd.AddCommand(getCommand)
}

func validateOutputFormat(outputFormat string) error {
	switch outputFormat {
	case "json", "yaml", "yml", "table", "default":
	default:
		return fmt.Errorf("invalid output format: %s", outputFormat)
	}
	return nil
}

func validateResourceType(resourceType string) (*k8s.ResourceInfo, error) {
	info, exists := k8s.ResourceRegistry[resourceType]
	if !exists {
		return nil, fmt.Errorf("invalid resource type: %s", resourceType)
	}
	return &info, nil
}
