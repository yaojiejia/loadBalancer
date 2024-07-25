package balancer

import (
	"fmt"
	"net/http"

	"github.com/yaojiejia/loadBalancer/proxy"
)

// baseLoadBalancer contains fields and methods common to both load balancers
type baseLoadBalancer struct {
	Port            string
	RoundRobinCount int
	Servers         []proxy.Server
}

// getNextAvailableServer is now part of baseLoadBalancer
func (lb *baseLoadBalancer) GetNextAvailableServer() proxy.Server {
	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	for !server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	}
	lb.RoundRobinCount++
	return server
}

// rrLoadBalancer for round-robin strategy
type rrLoadBalancer struct {
	baseLoadBalancer // Embedding
}

// New function for rrLoadBalancer
func NewRRLoadBalancer(port string, servers []proxy.Server) *rrLoadBalancer {
	return &rrLoadBalancer{
		baseLoadBalancer: baseLoadBalancer{
			Port:    port,
			Servers: servers,
		},
	}
}

// rrServeProxy uses the base getNextAvailableServer method
func (lb *rrLoadBalancer) RrServeProxy(res http.ResponseWriter, req *http.Request) {
	targetServer := lb.GetNextAvailableServer()
	fmt.Printf("forwarding request to address: %q\n", targetServer.Address())
	targetServer.Serve(res, req)
}
