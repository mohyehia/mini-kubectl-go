package k8s

type ResourceKind string

const (
	POD        ResourceKind = "Pod"
	SERVICE    ResourceKind = "Service"
	DEPLOYMENT ResourceKind = "Deployment"
	NAMESPACE  ResourceKind = "Namespace"
	NODE       ResourceKind = "Node"
)

// ResourceInfo stores the absolute blueprint needed to route a resource.
type ResourceInfo struct {
	GroupPath  string
	Namespaced bool
	PluralName string
	Kind       ResourceKind
}

type resourceSpec struct {
	info    ResourceInfo
	aliases []string
}

var ResourceRegistry = make(map[string]ResourceInfo)

func init() {
	resourceSpecs := []resourceSpec{
		{
			info: ResourceInfo{
				GroupPath:  "api/v1",
				Namespaced: true,
				PluralName: "pods",
				Kind:       POD,
			}, aliases: []string{"po", "pod", "pods"},
		},
		{
			info: ResourceInfo{
				GroupPath:  "api/v1",
				Namespaced: true,
				PluralName: "services",
				Kind:       SERVICE,
			}, aliases: []string{"svc", "service", "services"},
		},
		{
			info: ResourceInfo{
				GroupPath:  "apis/apps/v1",
				Namespaced: true,
				PluralName: "deployments",
				Kind:       DEPLOYMENT,
			}, aliases: []string{"deploy", "deployment", "deployments"},
		},
		{
			info: ResourceInfo{
				GroupPath:  "api/v1",
				Namespaced: false,
				PluralName: "namespaces",
				Kind:       NAMESPACE,
			}, aliases: []string{"ns", "namespace", "namespaces"},
		},
		{
			info: ResourceInfo{
				GroupPath:  "api/v1",
				Namespaced: false,
				PluralName: "nodes",
				Kind:       NODE,
			}, aliases: []string{"no", "node", "nodes"},
		},
	}

	for _, spec := range resourceSpecs {
		for _, alias := range spec.aliases {
			ResourceRegistry[alias] = spec.info
		}
	}
}
