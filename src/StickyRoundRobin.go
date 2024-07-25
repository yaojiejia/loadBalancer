package main

import (
	"fmt"
	"net/http"
)

// srrLoadBalancer for sticky round-robin strategy
type srrLoadBalancer struct {
	baseLoadBalancer // Embedding
	userMap          map[string]Server
}

// New function for srrLoadBalancer
func newSRRLoadBalancer(port string, servers []Server) *srrLoadBalancer {
	return &srrLoadBalancer{
		baseLoadBalancer: baseLoadBalancer{
			port:    port,
			servers: servers,
		},
		userMap: make(map[string]Server),
	}
}

// getNextAvailableServer is overridden to include sticky session logic
func (lb *srrLoadBalancer) getNextAvailableServer(ip string) Server {
	if server, exists := lb.userMap[ip]; exists {
		return server
	}
	server := lb.baseLoadBalancer.getNextAvailableServer()
	lb.userMap[ip] = server
	return server
}

// srrServeProxy uses the overridden getNextAvailableServer method
func (lb *srrLoadBalancer) srrServeProxy(res http.ResponseWriter, req *http.Request, ip string) {
	targetServer := lb.getNextAvailableServer(ip)
	fmt.Printf("forwarding request to address: %q\n", targetServer.Address())
	if targetServer.IsAlive() {
		fmt.Println("Server is alive")
	} else {
		fmt.Println("Server is down")
	}
	targetServer.Serve(res, req)
}
