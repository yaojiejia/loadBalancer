package balancer

import (
	"fmt"
	"hash/fnv"
	"net/http"

	"github.com/yaojiejia/loadBalancer/proxy"
)

type IPLoadBalancer struct {
	Port    string
	Servers []proxy.Server
}

func NewIPLoadBalancer(port string, servers []proxy.Server) *IPLoadBalancer {
	return &IPLoadBalancer{
		Port:    port,
		Servers: servers,
	}
}

func (lb *IPLoadBalancer) GetNextAvailableServer(ip string) proxy.Server {
	hash := fnv.New32a()
	hash.Write([]byte(ip))
	index := int(int(hash.Sum32()) % len(lb.Servers))
	fmt.Println(hash)
	return lb.Servers[index]
}

func (lb *IPLoadBalancer) IPServeProxy(res http.ResponseWriter, req *http.Request, ip string) {
	targetServer := lb.GetNextAvailableServer(ip)
	fmt.Printf("forwarding request to address: %q\n", targetServer.Address())

	targetServer.Serve(res, req)
}
