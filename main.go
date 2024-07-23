package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(res http.ResponseWriter, req *http.Request)
}

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

type loadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func NewloadBalancer(port string, servers []Server) *loadBalancer {
	return &loadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}

}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func (s *simpleServer) Address() string {
	return s.addr
}

func (s *simpleServer) IsAlive() bool {
	return true
}
func (s *simpleServer) Serve(res http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(res, req)
}

func (lb *loadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}

	lb.roundRobinCount++
	return server
}

func (lb *loadBalancer) serveProxy(res http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address: %q\n", targetServer.Address())
	targetServer.Serve(res, req)
}

func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.go.dev"),
		newSimpleServer("https://www.spotify.com"),
	}

	lb := NewloadBalancer("8800", servers)
	handleRedirect := func(res http.ResponseWriter, req *http.Request) {
		lb.serveProxy(res, req)
	}
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Program started at %s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)

}
