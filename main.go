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
	// Configures http server.

	var s http.Server
	flag.BoolVar(&http2.VerboseLogs, "verbose", false, "verbose HTTP/2 debugging.")
	flag.Parse()
	s.Addr = *addr

	// Opens database connection and retrieves list of posts to store in memory
	p, err := database.Posts()
	if err != nil {
		log.Fatal(err)
	}

	registerHandlers(p)

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
