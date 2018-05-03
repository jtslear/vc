package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"time"
)

func getTLS(host string) (valid bool, cert []*x509.Certificate, reason string) {
	conn, err := tls.Dial("tcp", host, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		cert = conn.ConnectionState().PeerCertificates
		return false, cert, err.Error()
	}
	conn.Close()
	cert = conn.ConnectionState().PeerCertificates
	return true, cert, ""
}

func expired(i time.Time) bool {
	if time.Now().Before(i) {
		return false
	}
	return true
}

func main() {
	flag.Parse()
	for _, v := range flag.Args() {
		ok, certs, msg := getTLS(v)
		var cert = &x509.Certificate{}
		cert = certs[0]
		if !ok {
			fmt.Printf("%s: notok: %s\n", v, msg)
		} else {
			if expired(cert.NotAfter) {
				fmt.Printf("%s: notok - expired: %s\n", v, msg)
			} else {
				fmt.Printf("%s: ok - %s\n", v, cert.NotAfter)
			}
		}
	}
}
