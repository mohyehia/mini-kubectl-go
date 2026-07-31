/*
Package cmd Copyright © 2026 mohyehia <mohammedyehia99@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/mohyehia/mini-kubectl/pkg/k8s"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mini-kubectl",
	Short: "kubectl controls the Kubernetes cluster manager.",
	Long: `kubectl controls the Kubernetes cluster manager.

 Find more information at: https://kubernetes.io/docs/reference/kubectl/

Basic Commands:
  get             Display one or many resources
  create          Create a resource from a file or from stdin
  apply           Apply a configuration to a resource by file name or stdin
  delete          Delete resources by file names, stdin, resources and names, or by resources and label selector

Troubleshooting and Debugging Commands:
  describe        Show details of a specific resource or group of resources
  logs            Print the logs for a container in a pod

Deploy Commands:
  rollout         Manage the rollout of a resource

Usage:
  kubectl [flags] [options]

Use "kubectl <command> --help" for more information about a given command.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var kubeConfigDefaultPath = "~/.kube/config"

func init() {
	rootCmd.PersistentFlags().StringVarP(&appState.configFile, "config", "c", kubeConfigDefaultPath, "config file (default is $HOME/.kube/config)")
	rootCmd.PersistentFlags().StringVarP(&appState.namespace, "namespace", "n", "default", "namespace to use")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {

		configFile, err := resolveConfigPath(appState.configFile)
		if err != nil {
			return fmt.Errorf("error resolving config path %q: %w", configFile, err)
		}
		appState.configFile = configFile

		bytes, err := os.ReadFile(appState.configFile)
		if err != nil {
			return fmt.Errorf("error reading config file: %w", err)
		}

		err = yaml.Unmarshal(bytes, &appState.kubeConfig)
		if err != nil {
			return fmt.Errorf("error parsing config file: %w", err)
		}

		appState.currentClusterConnection, err = currentClusterConnectionFromConfig(appState.kubeConfig)

		if err != nil {
			return fmt.Errorf("failed to get current cluster connection from config: %w", err)
		}

		appState.k8sClient, err = k8s.NewClient(appState.currentClusterConnection)
		if err != nil {
			return fmt.Errorf("failed to initialize k8s client: %w", err)
		}
		return nil
	}
}

func resolveConfigPath(path string) (string, error) {
	if strings.Contains(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return path, err
		}
		path = strings.Replace(path, "~", homeDir, 1)
	}
	return path, nil
}

func currentClusterConnectionFromConfig(config k8s.KubeConfig) (k8s.CurrentClusterConnection, error) {
	currentClusterConnection := k8s.CurrentClusterConnection{}
	// Resolve the active context before looking up cluster and user details.
	for i := range config.Contexts {
		if config.Contexts[i].Name == config.CurrentContext {
			currentClusterConnection.CurrentClusterName = config.Contexts[i].Context.Cluster
			currentClusterConnection.CurrentUserName = config.Contexts[i].Context.User
		}
	}
	if strings.TrimSpace(currentClusterConnection.CurrentClusterName) == "" || strings.TrimSpace(currentClusterConnection.CurrentUserName) == "" {
		return currentClusterConnection, fmt.Errorf("current cluster name or user name not found in kubeconfig")
	}

	for i := range config.Clusters {
		if config.Clusters[i].Name == currentClusterConnection.CurrentClusterName {
			currentClusterConnection.ServerURL = config.Clusters[i].Cluster.Server
			currentClusterConnection.CertificateAuthorityData = config.Clusters[i].Cluster.CertificateAuthorityData
		}
	}
	if strings.TrimSpace(currentClusterConnection.ServerURL) == "" || strings.TrimSpace(currentClusterConnection.CertificateAuthorityData) == "" {
		return currentClusterConnection, fmt.Errorf("current cluster server URL or certificate authority data not found in kubeconfig")
	}

	for i := range config.Users {
		if config.Users[i].Name == currentClusterConnection.CurrentUserName {
			currentClusterConnection.ClientCertificateData = config.Users[i].User.ClientCertificateData
			currentClusterConnection.ClientKeyData = config.Users[i].User.ClientKeyData
		}
	}
	if strings.TrimSpace(currentClusterConnection.ClientCertificateData) == "" || strings.TrimSpace(currentClusterConnection.ClientKeyData) == "" {
		return currentClusterConnection, fmt.Errorf("current user client certificate data or client key data not found in kubeconfig")
	}

	return currentClusterConnection, nil
}
