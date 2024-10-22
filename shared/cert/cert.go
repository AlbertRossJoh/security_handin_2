package cert

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc/credentials"
)

func LoadTLSServerCredentials(certPath string, keyPath string, caCertPath string) (credentials.TransportCredentials, error) {
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
		RootCAs:            certPool,
		InsecureSkipVerify: false,
	}

	return credentials.NewTLS(config), nil
}

func LoadTLSClientCredentials(certPath string, keyPath string, caCertPath string, dockerId string) credentials.TransportCredentials {
	clientCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Panicln("Could not load client certificate")
	}

	certPool, err := LoadCaCertPool(caCertPath)
	if err != nil {
		log.Panicln("Could not load CA cert")
	}
	config := &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          certPool,
		RootCAs:            certPool,
		ServerName:         dockerId,
		InsecureSkipVerify: false,
	}

	return credentials.NewTLS(config)
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
