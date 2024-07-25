package main

import (
	"errors"
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"
)

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) (*simpleServer, error) {
	serverUrl, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}, nil
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Print("ERROR HERE! CHECK SIMPLESERVER")
		os.Exit(1)
	}
}

func validateURL(addr string) error {
	parsed, err := url.ParseRequestURI(addr)
	if err != nil {
		return err
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return errors.New("unsupported protocol scheme (needs to start in http/https)")
	}

	return nil
}
