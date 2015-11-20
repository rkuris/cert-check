package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"time"
)

type SiteCheck struct {
	notify int64
	sites []string
}

func main() {
	check := new(SiteCheck)
	parseArguments( check )

	config := tls.Config{InsecureSkipVerify: false}
	for _, site := range check.sites {
		dosite(site, check.notify, config)
	}
}

func parseArguments( check *SiteCheck ) {
	flag.Int64Var( &check.notify, "notify", 220, "Notify if fewer than this number of days are presented" )
	flag.Parse()
	check.sites = flag.Args()
}

func dosite(site string, notify int64,  config tls.Config) {
	conn, err := tls.Dial("tcp", site, &config)
	if err != nil {
		log.Fatalf("client(%s): %s", site, err)
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
