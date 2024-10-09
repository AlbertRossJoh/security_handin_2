package cert

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	"google.golang.org/grpc/credentials"
)

func LoadTLSCredentials(certPath string, keyPath string, caCertPath string) (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	certPool, err := LoadCaCertPool(caCertPath)
	if err != nil {
		return nil, err
	}
	// Create the credentials and return it
	config := &tls.Config{
		Certificates:       []tls.Certificate{serverCert},
		ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          certPool,
		InsecureSkipVerify: true,
		//RootCAs:            certPool,
	}

	return credentials.NewTLS(config), nil
}

func LoadCaCertPool(caCertPath string) (*x509.CertPool, error) {
	pemServerCA, err := os.ReadFile(caCertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}
	return certPool, nil
}
