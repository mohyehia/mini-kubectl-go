package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the client and server version information for the current context.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Client Version: ", "v1.0.0")
		fmt.Println("Server Version: ", "v1.35.5+k3s1")
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
