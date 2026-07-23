package k8s

type CurrentClusterConnection struct {
	CurrentClusterName       string
	CurrentUserName          string
	ServerURL                string
	CertificateAuthorityData string
	ClientCertificateData    string
	ClientKeyData            string
}
