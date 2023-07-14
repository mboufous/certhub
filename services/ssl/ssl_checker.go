package ssl

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"net"
	"time"
)

// CertificateStatus represents the status of an SSL certificate.
type CertificateStatus string

const (
	CertificateStatusValid        CertificateStatus = "valid"
	CertificateStatusExpired      CertificateStatus = "expired"
	CertificateStatusInvalid      CertificateStatus = "invalid"
	CertificateStatusExpiringSoon CertificateStatus = "expiringSoon"
)

// CertificateInfo contains the SSL certificate information.
type CertificateInfo struct {
	Subject      string
	Issuer       string
	Expires      time.Time
	Status       CertificateStatus
	SerialNumber string
	DNSNames     []string
	IPAddresses  []net.IP
}

// GetCertificateInfo returns the SSL certificate information for a given domain.
func GetCertificateInfo(domain string) (*CertificateInfo, error) {
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to establish TLS connection: %v", err)
	}
	defer conn.Close()

	// Get the peer certificate from the TLS connection
	cert := conn.ConnectionState().PeerCertificates[0]
	certInfo := &CertificateInfo{
		Subject:      cert.Subject.CommonName,
		Issuer:       cert.Issuer.CommonName,
		Expires:      cert.NotAfter,
		Status:       getCertificateStatus(cert),
		SerialNumber: hex.EncodeToString(cert.SerialNumber.Bytes()),
		DNSNames:     cert.DNSNames,
		IPAddresses:  cert.IPAddresses,
	}

	return certInfo, nil
}

// getCertificateStatus returns the status of an SSL certificate.
func getCertificateStatus(cert *x509.Certificate) CertificateStatus {
	currentTime := time.Now()

	if currentTime.After(cert.NotAfter) {
		return CertificateStatusExpired
	} else if currentTime.Before(cert.NotBefore) {
		return CertificateStatusInvalid
	}

	return CertificateStatusValid
}
