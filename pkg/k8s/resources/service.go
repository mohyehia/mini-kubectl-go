package resources

type Service struct {
	Metadata ObjectMetadata `json:"metadata"`
	Spec     ServiceSpec    `json:"spec"`
	Status   ServiceStatus  `json:"status"`
}

type ServiceSpec struct {
	Ports        []ServicePort `json:"ports"`
	ClusterIP    string        `json:"clusterIP"`
	Type         string        `json:"type"`
	ExternalIPs  []string      `json:"externalIPs"`
	ExternalName string        `json:"externalName"`
}

type ServicePort struct {
	Name       string `json:"name"`
	Protocol   string `json:"protocol"`
	Port       int    `json:"port"`
	TargetPort int    `json:"targetPort"`
}

type ServiceStatus struct {
	LoadBalancer LoadBalancerStatus `json:"loadBalancer"`
}

type LoadBalancerStatus struct {
	Ingress []LoadBalancerIngress `json:"ingress"`
}

type LoadBalancerIngress struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
}
