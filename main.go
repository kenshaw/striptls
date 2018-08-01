// Command striptls provides a simple HTTP proxy that strips tls.
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	flagListen   = flag.String("l", "localhost:80", "listen address")
	flagRemote   = flag.String("r", "https://localhost:443", "remote address")
	flagInsecure = flag.Bool("i", true, "insecure")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	// parse url
	u, err := url.Parse(*flagRemote)
	if err != nil {
		return err
	}

	// override default transport
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
		InsecureSkipVerify: *flagInsecure,
	}

	// proxy
	mux := http.NewServeMux()
	mux.Handle("/", httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: u.Scheme,
		Host:   u.Host,
	}))
	return http.ListenAndServe(*flagListen, mux)
}
