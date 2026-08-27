package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mohyehia/mini-kubectl/pkg/k8s"
	"github.com/spf13/cobra"
)

var (
	waitFlag    bool
	timeoutFlag time.Duration
)

var deleteCommand = &cobra.Command{
	Use:   "delete [TYPE] [NAME]",
	Short: "Delete resources by type and name.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		resourceType := args[0]
		resourceName := args[1]

		// Validate resource type and name
		info, err := validateResourceType(resourceType)
		if err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), timeoutFlag)
		defer cancel()

		fmt.Printf("Deleting resource: Type=%s, Name=%s\n", info.Kind, resourceName)
		_, err = appState.k8sClient.DeleteResource(ctx, info.Kind, appState.namespace, resourceName)
		if err != nil {
			return err
		}
		// Handle wait flag if needed
		if waitFlag {
			fmt.Printf("Waiting for resource %s/%s to be fully deleted...\n", info.Kind, resourceName)
			if err := k8s.WaitForDeletion(ctx, os.Stdout, appState.k8sClient, info.Kind, appState.namespace, resourceName, 500*time.Millisecond); err != nil {
				return err
			}
		}
		fmt.Printf("%s %q deleted\n", info.Kind, resourceName)
		return nil
	},
}

func init() {
	deleteCommand.Flags().BoolVar(&waitFlag, "wait", true, "wait for resources to be fully deleted from cluster")
	deleteCommand.Flags().DurationVar(&timeoutFlag, "timeout", 30*time.Second, "the length of time to wait before giving up on a delete operation")
	rootCmd.AddCommand(deleteCommand)
}
