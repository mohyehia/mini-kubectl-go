package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
)

var followFlag bool
var logsCommand = &cobra.Command{
	Use:   "logs [POD_NAME]",
	Short: "Print or stream the logs for a container in a pod",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		podName := args[0]
		// Validate pod name
		if strings.TrimSpace(podName) == "" {
			return fmt.Errorf("pod name cannot be empty")
		}
		ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
		defer stop()

		err := appState.k8sClient.StreamLogs(ctx, os.Stdout, appState.namespace, podName, followFlag)

		if err != nil && ctx.Err() != nil {
			fmt.Printf("log streaming interrupted by user for pod %s", podName)
			return nil
		}
		return nil
	},
}

func init() {
	logsCommand.Flags().BoolVarP(&followFlag, "follow", "f", false, "Specify if the logs should be streamed")
	rootCmd.AddCommand(logsCommand)
}
