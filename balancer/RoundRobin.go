package balancer

import (
	"fmt"
	"net/http"

	"github.com/yaojiejia/loadBalancer/proxy"
)

// baseLoadBalancer contains fields and methods common to both load balancers
type BaseLoadBalancer struct {
	Port            string
	RoundRobinCount int
	Servers         []proxy.Server
}

// getNextAvailableServer is now part of baseLoadBalancer
func (lb *BaseLoadBalancer) GetNextAvailableServer() proxy.Server {
	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	for !server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	}
	lb.RoundRobinCount++
	return server
}

// rrLoadBalancer for round-robin strategy
type RRLoadBalancer struct {
	BaseLoadBalancer // Embedding
}

// New function for rrLoadBalancer
func NewRRLoadBalancer(port string, servers []proxy.Server) *RRLoadBalancer {
	return &RRLoadBalancer{
		BaseLoadBalancer: BaseLoadBalancer{
			Port:    port,
			Servers: servers,
		},
	}
}

// rrServeProxy uses the base getNextAvailableServer method
func (lb *RRLoadBalancer) RrServeProxy(res http.ResponseWriter, req *http.Request, ip string) {
	targetServer := lb.GetNextAvailableServer()
	fmt.Printf("forwarding request to address: %q\n", targetServer.Address())
	targetServer.Serve(res, req)
}
