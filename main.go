package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/bradfitz/http2"
)

var (
	addr     = flag.String("tls", "localhost:4430", "TLS address to listen on")
	httpAddr = flag.String("http", "", "address to listen for regular http on")
	prod     = flag.Bool("prod", false, "whether to configure for production")
)

var p []Post = samplePost()

func main() {
	var s http.Server
	flag.BoolVar(&http2.VerboseLogs, "verbose", false, "verbose HTTP/2 debugging.")
	flag.Parse()
	s.Addr = *addr

	registerHandlers()

	url := "https://" + *addr + "/"
	log.Printf("Listening on " + url)
	http2.ConfigureServer(&s, &http2.Server{})

	if *httpAddr != "" {
		go func() { log.Fatal(http.ListenAndServe(*httpAddr, nil)) }()
	}

	go func() {
		log.Fatal(s.ListenAndServeTLS("out/dev.crt", "out/dev.key"))
	}()

	select {}
}
