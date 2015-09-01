package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/billychappell/downloader/database"
	"github.com/bradfitz/http2"
)

var (
	addr     = flag.String("tls", "localhost:4430", "TLS address to listen on")
	httpAddr = flag.String("http", "", "address to listen for regular http on")
	prod     = flag.Bool("prod", false, "whether to configure for production")
)

func main() {

	// Initialize server and parse flags.
	var s http.Server
	flag.BoolVar(&http2.VerboseLogs, "verbose", false, "verbose HTTP/2 debugging")
	flag.Parse()
	s.Addr = *addr

	// Open database and retrieve posts.
	p, err := database.GetPosts()
	if err != nil {
		log.Fatal(err)
	}

	// Registers handlers and logs URL
	// and port, configured for HTTP/2.
	registerHandlers(p)
	url := "https://" + *addr + "/"
	log.Printf("Listening on " + url)
	http2.ConfigureServer(&s, &http2.Server{})

	// Starts listening over regular HTTP
	// if the "-http" flag is non-empty.
	if *httpAddr != "" {
		go func() { log.Fatal(http.ListenAndServe(*httpAddr, nil)) }()
	}

	// Concurrently begins listening at URL for
	// tls encrypted HTTP/2 communications.
	go func() {
		log.Fatal(s.ListenAndServeTLS("out/dev.crt", "out/dev.key"))
	}()

	// Blocks forever.
	select {}
}
