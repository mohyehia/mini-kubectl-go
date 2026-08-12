package k8s

func getURL(serverURL string, kind ResourceKind, namespace string) string {
	switch kind {
	case NAMESPACE:
		return serverURL + "/api/v1/namespaces"
	case NODE:
		return serverURL + "/api/v1/nodes"
	case POD:
		return serverURL + "/api/v1/namespaces/" + namespace + "/pods"
	case SERVICE:
		return serverURL + "/api/v1/namespaces/" + namespace + "/services"
	case DEPLOYMENT:
		return serverURL + "/apis/apps/v1/namespaces/" + namespace + "/deployments"
	default:
		return ""
	}
}
