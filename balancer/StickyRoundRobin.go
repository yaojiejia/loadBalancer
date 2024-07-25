package balancer

import (
	"fmt"
	"net/http"

	"github.com/yaojiejia/loadBalancer/proxy"
)

// srrLoadBalancer for sticky round-robin strategy
type srrLoadBalancer struct {
	baseLoadBalancer // Embedding
	UserMap          map[string]proxy.Server
}

// New function for srrLoadBalancer
func NewSRRLoadBalancer(port string, servers []proxy.Server) *srrLoadBalancer {
	return &srrLoadBalancer{
		baseLoadBalancer: baseLoadBalancer{
			Port:    port,
			Servers: servers,
		},
		UserMap: make(map[string]proxy.Server),
	}
}

// getNextAvailableServer is overridden to include sticky session logic
func (lb *srrLoadBalancer) GetNextAvailableServer(ip string) proxy.Server {
	if server, exists := lb.UserMap[ip]; exists {
		return server
	}
	server := lb.baseLoadBalancer.GetNextAvailableServer()
	lb.UserMap[ip] = server
	return server
}

// srrServeProxy uses the overridden getNextAvailableServer method
func (lb *srrLoadBalancer) SrrServeProxy(res http.ResponseWriter, req *http.Request, ip string) {
	targetServer := lb.GetNextAvailableServer(ip)
	fmt.Printf("forwarding request to address: %q\n", targetServer.Address())
	if targetServer.IsAlive() {
		fmt.Println("Server is alive")
	} else {
		fmt.Println("Server is down")
	}
	targetServer.Serve(res, req)
}
