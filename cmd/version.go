package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version will be set at build time via -ldflags. Defaults to dev if compiled directly.
var Version = "v0.1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the client and server version information for the current context.",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Client Version: %s\n", Version)

		if appState.k8sClient == nil {
			fmt.Println("Server info is not available at the moment.")
			return nil
		}

		serverVersion, err := appState.k8sClient.GetServerVersion()
		if err != nil {
			fmt.Printf("Server Version: unavailable (%v)\n", err)
			return nil
		}
		fmt.Printf("Server Version: %s (Platform: %s)\n", serverVersion.GitVersion, serverVersion.Platform)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
