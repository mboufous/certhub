package models

import (
	"net"
	"time"
)

type User struct {
	Id             int
	Name           string
	Email          string
	Password       string
	CreatedAt      time.Time
	LastLogin      time.Time
	Disabled       bool
	ManagedDomains []Domain
}

type Domain struct {
	Id              int
	Name            string
	CreatedAt       time.Time
	Active          bool
	CertificateInfo *CertificateInfo
}

type CertificateStatus string

const (
	CertificateStatusValid   CertificateStatus = "Valid"
	CertificateStatusExpired CertificateStatus = "Expired"
	CertificateStatusInvalid CertificateStatus = "Invalid"
)

type CertificateInfo struct {
	Subject      string
	Issuer       string
	Expires      time.Time
	Status       CertificateStatus
	SerialNumber string
	DNSNames     []string
	IPAddresses  []net.IP
}
