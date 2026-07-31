package cmd

import "github.com/mohyehia/mini-kubectl/pkg/k8s"

var appState struct {
	configFile               string
	namespace                string
	kubeConfig               k8s.KubeConfig
	currentClusterConnection k8s.CurrentClusterConnection
	k8sClient                *k8s.Client
}
