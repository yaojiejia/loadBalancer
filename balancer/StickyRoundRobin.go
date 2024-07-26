package balancer

import (
	"fmt"
	"net/http"

	"github.com/yaojiejia/loadBalancer/proxy"
)

// srrLoadBalancer for sticky round-robin strategy
type SRRLoadBalancer struct {
	BaseLoadBalancer // Embedding
	UserMap          map[string]proxy.Server
}

// New function for srrLoadBalancer
func NewSRRLoadBalancer(port string, servers []proxy.Server) *SRRLoadBalancer {
	return &SRRLoadBalancer{
		BaseLoadBalancer: BaseLoadBalancer{
			Port:    port,
			Servers: servers,
		},
		UserMap: make(map[string]proxy.Server),
	}
}

// getNextAvailableServer is overridden to include sticky session logic
func (lb *SRRLoadBalancer) GetNextAvailableServer(ip string) proxy.Server {
	if server, exists := lb.UserMap[ip]; exists {
		return server
	}
	server := lb.BaseLoadBalancer.GetNextAvailableServer()
	lb.UserMap[ip] = server
	return server
}

// srrServeProxy uses the overridden getNextAvailableServer method
func (lb *SRRLoadBalancer) SrrServeProxy(res http.ResponseWriter, req *http.Request, ip string) {
	targetServer := lb.GetNextAvailableServer(ip)
	fmt.Printf("forwarding request to address: %q\n", targetServer.Address())
	if targetServer.IsAlive() {
		fmt.Println("Server is alive")
	} else {
		fmt.Println("Server is down")
	}
	targetServer.Serve(res, req)
}
