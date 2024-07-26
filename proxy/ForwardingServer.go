package proxy

import (
	"errors"
	"fmt"
	"net/http/httputil"
	"net/url"
	"os"
)

type SimpleServer struct {
	Addr  string
	Proxy *httputil.ReverseProxy
}

func NewSimpleServer(addr string) (*SimpleServer, error) {
	serverUrl, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	return &SimpleServer{
		Addr:  addr,
		Proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}, nil
}

func HandleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		fmt.Print("ERROR HERE! CHECK SIMPLESERVER")
		os.Exit(1)
	}
}

func ValidateURL(addr string) error {
	parsed, err := url.ParseRequestURI(addr)
	if err != nil {
		return err
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return errors.New("unsupported protocol scheme (needs to start in http/https)")
	}

	return nil
}
