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

var configFile string
var namespace string
var kubeConfigDefaultPath = "~/.kube/config"
var KubeConfig k8s.KubeConfig
var currentClusterConnection k8s.CurrentClusterConnection
var K8sClient *k8s.Client

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", kubeConfigDefaultPath, "config file (default is $HOME/.kube/config)")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "namespace to use")

	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {

		if strings.Contains(configFile, "~") {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				panic(err)
			}
			configFile = strings.Replace(configFile, "~", homeDir, 1)
		}

		bytes, err := os.ReadFile(configFile)
		if err != nil {
			fmt.Printf("Error reading config file: %s\n", err)
			os.Exit(1)
		}

		err = yaml.Unmarshal(bytes, &KubeConfig)
		if err != nil {
			fmt.Printf("Error parsing config file: %s\n", err)
			os.Exit(1)
		}

		currentClusterConnection = k8s.CurrentClusterConnection{}

		// Important step: we need to find the current context first inorder for the remaining parts to be correct
		for i := range KubeConfig.Contexts {
			if KubeConfig.Contexts[i].Name == KubeConfig.CurrentContext {
				currentClusterConnection.CurrentClusterName = KubeConfig.Contexts[i].Context.Cluster
				currentClusterConnection.CurrentUserName = KubeConfig.Contexts[i].Context.User
			}
		}

		for i := range KubeConfig.Clusters {
			if KubeConfig.Clusters[i].Name == currentClusterConnection.CurrentClusterName {
				currentClusterConnection.ServerURL = KubeConfig.Clusters[i].Cluster.Server
				currentClusterConnection.CertificateAuthorityData = KubeConfig.Clusters[i].Cluster.CertificateAuthorityData
			}
		}

		for i := range KubeConfig.Users {
			if KubeConfig.Users[i].Name == currentClusterConnection.CurrentUserName {
				currentClusterConnection.ClientCertificateData = KubeConfig.Users[i].User.ClientCertificateData
				currentClusterConnection.ClientKeyData = KubeConfig.Users[i].User.ClientKeyData
			}
		}

		//fmt.Printf("CurrentClusterName: %s\n", currentClusterConnection.CurrentClusterName)
		//fmt.Printf("CurrentUserName: %s\n", currentClusterConnection.CurrentUserName)
		//fmt.Printf("ServerURL: %s\n", currentClusterConnection.ServerURL)
		//fmt.Printf("CertificateAuthorityData: %s\n", currentClusterConnection.CertificateAuthorityData)
		//fmt.Printf("ClientCertificateData: %s\n", currentClusterConnection.ClientCertificateData)
		//fmt.Printf("ClientKeyData: %s\n", currentClusterConnection.ClientKeyData)

		K8sClient, err = k8s.NewClient(currentClusterConnection)
		if err != nil {
			return fmt.Errorf("failed to initialize k8s client: %w", err)
		}
		return nil
	}
}
