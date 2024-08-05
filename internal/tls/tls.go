package tls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

// LoadServerTLS loads TLS certificates for the server
func LoadServerTLS() (*tls.Config, error) {
	cert, err := tls.LoadX509KeyPair("certs/server.crt", "certs/server.key")
	if err != nil {
		return nil, fmt.Errorf("failed to load server certificates: %v", err)
	}

	// Create a certificate pool from the CA
    caCert, err := os.ReadFile("certs/ca.crt")
    if err != nil {
        return nil, fmt.Errorf("failed to read CA cert: %v", err)
    }
    caPool := x509.NewCertPool()
    caPool.AppendCertsFromPEM(caCert)

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
        ClientCAs:    caPool,
	}, nil
}

// LoadClientTLS loads TLS configuration for the client
func LoadClientTLS() (*tls.Config, error) {
	caCert, err := os.ReadFile("certs/ca.crt")
	if err != nil {
		return nil, fmt.Errorf("failed to read CA certificate: %v", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caCert) {
		return nil, fmt.Errorf("failed to add CA certificate to pool")
	}

	return &tls.Config{
		RootCAs: certPool,
		InsecureSkipVerify: false,
	}, nil
}
