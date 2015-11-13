package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"time"
)

const notify int64 = 220

func main() {
	sites := []string{
		"okta.okta.com:443",
		"authbank.com:443",
		"uber.com:443",
	}

	config := tls.Config{InsecureSkipVerify: false}
	for _, site := range sites {
		dosite(site, config)
	}
}
func dosite(site string, config tls.Config) {
	conn, err := tls.Dial("tcp", site, &config)
	if err != nil {
		log.Fatalf("client: %s", err)
	}
	defer conn.Close()
	state := conn.ConnectionState()
	for pos, cert := range state.PeerCertificates {
		{
			diff := cert.NotAfter.Unix() - time.Now().Unix()
			diff /= 86400
			certkind := "Certificate"
			if pos > 0 {
				certkind = "Certificate authority"
			}
			if diff < notify {
				fmt.Printf("%s %s expires in %d days at %s (warning at %d days)\n", site, certkind, diff, cert.NotAfter, notify)
			}
		}
	}
}
