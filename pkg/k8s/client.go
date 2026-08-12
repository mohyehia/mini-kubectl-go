package k8s

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ClusterOperations interface {
	GetServerVersion() (*ServerVersion, error)

	GetResourceList(kind ResourceKind, namespace string) ([]byte, error)
}

type Client struct {
	HttpClient *http.Client
	ServerURL  string
}

// NewClient This implementation currently supports mTLS kubeconfigs only.
func NewClient(connection CurrentClusterConnection) (*Client, error) {
	// 1. Decode Base64 Certificate Authority (CA)
	caBytes, err := base64.StdEncoding.DecodeString(connection.CertificateAuthorityData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode certificate authority data: %w", err)
	}

	// 2. Decode Base64 Client Certificate & Key
	clientCertificateBytes, err := base64.StdEncoding.DecodeString(connection.ClientCertificateData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode client certificate data: %w", err)
	}

	clientKeyBytes, err := base64.StdEncoding.DecodeString(connection.ClientKeyData)
	if err != nil {
		return nil, fmt.Errorf("failed to decode client key data: %w", err)
	}

	// 3. Build Root CA Pool (Client -> Server Trust)
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(caBytes) {
		return nil, fmt.Errorf("failed to append CA certificate to pool")
	}

	// 4. Build KeyPair (Server -> Client Trust)
	cert, err := tls.X509KeyPair(clientCertificateBytes, clientKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to load client certificate and key pair: %w", err)
	}

	// 5. Construct TLS Configuration
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	return &Client{
		HttpClient: &http.Client{
			Transport: transport,
			Timeout:   5 * time.Second,
		},
		ServerURL: connection.ServerURL,
	}, nil
}

func (c *Client) GetServerVersion() (*ServerVersion, error) {
	url := c.ServerURL + "/version"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the server: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned unexpected status code: %d", resp.StatusCode)
	}

	var serverVersion ServerVersion
	if err := json.NewDecoder(resp.Body).Decode(&serverVersion); err != nil {
		return nil, fmt.Errorf("failed to unmarshal server version: %w", err)
	}
	return &serverVersion, nil
}

func (c *Client) GetResourceList(kind ResourceKind, namespace string) ([]byte, error) {
	url := getURL(c.ServerURL, kind, namespace)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the server: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("server returned unexpected status code: %d", resp.StatusCode)
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	return bytes, nil
}
