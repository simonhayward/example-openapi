package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/simonhayward/example-openapi/m/api"
)

func main() {
	port := flag.String("port", "8000", "Port http server")
	flag.Parse()

	server := api.NewServer()

	r := http.NewServeMux()

	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    net.JoinHostPort("0.0.0.0", *port),
	}

	log.Printf("serving on %s", *port)
	log.Fatal(s.ListenAndServe())
}
