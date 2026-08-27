package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var describeCommand = &cobra.Command{
	Use:   "describe (TYPE [NAME] | TYPE/NAME)",
	Short: "Show details of a specific resource or group of resources",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		resourceType := args[0]
		var resourceName string

		if len(args) == 2 {
			resourceName = args[1]
		}

		// Validate resource type and name
		info, err := validateResourceType(resourceType)
		if err != nil {
			return err
		}

		fmt.Printf("Describing resource (WIP): Type=%s, Name=%s\n", info.Kind, resourceName)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(describeCommand)
}
