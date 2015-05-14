package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/drone/config"

	"github.com/crewjam/awsconsoleauth"
)

func main() {
	listenAddress := flag.String("listen", ":8080", "The address the web server should listen on")
	configFile := flag.String("config", "awsauthd.conf", "The path to the configuration file")
	flag.Parse()

	config.SetPrefix("AWSAUTHD_")
	if err := config.Parse(*configFile); err != nil {
		log.Fatal(err)
	}

	if err := awsconsoleauth.Initialize(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Listening on %s\n", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
