package main

import (
	"fmt"
	"net/http"
)

// baseLoadBalancer contains fields and methods common to both load balancers
type baseLoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

// getNextAvailableServer is now part of baseLoadBalancer
func (lb *baseLoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

// rrLoadBalancer for round-robin strategy
type rrLoadBalancer struct {
	baseLoadBalancer // Embedding
}

// New function for rrLoadBalancer
func newRRLoadBalancer(port string, servers []Server) *rrLoadBalancer {
	return &rrLoadBalancer{
		baseLoadBalancer: baseLoadBalancer{
			port:    port,
			servers: servers,
		},
	}
}

// rrServeProxy uses the base getNextAvailableServer method
func (lb *rrLoadBalancer) rrServeProxy(res http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address: %q\n", targetServer.Address())
	targetServer.Serve(res, req)
}
