package main

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/google/go-containerregistry/pkg/registry"
	"github.com/spf13/pflag"
)

const (
	certPath = "/certs/tls.crt"
	keyPath  = "/certs/tls.key"
)

func main() {
	var (
		registryAddr string
	)
	flag.StringVar(&registryAddr, "registry-address", ":12345", "The address the registry binds to.")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	s := &http.Server{
		Addr:         registryAddr,
		Handler:      registry.New(),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	err := s.ListenAndServeTLS(certPath, keyPath)
	if err != nil {
		log.Fatalf("failed to start image registry: %s", err.Error())
	}

	defer s.Close()
}
