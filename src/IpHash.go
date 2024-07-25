package main

import (
	"fmt"
	"hash/fnv"
	"net/http"
)

type IPLoadBalancer struct {
	port    string
	servers []Server
}

func newIPLoadBalancer(port string, servers []Server) *IPLoadBalancer {
	return &IPLoadBalancer{
		port:    port,
		servers: servers,
	}
}

func (lb *IPLoadBalancer) getNextAvailableServer(ip string) Server {
	hash := fnv.New32a()
	hash.Write([]byte(ip))
	index := int(int(hash.Sum32()) % len(lb.servers))
	fmt.Println(hash)
	return lb.servers[index]
}

func (lb *IPLoadBalancer) IPServeProxy(res http.ResponseWriter, req *http.Request, ip string) {
	targetServer := lb.getNextAvailableServer(ip)
	fmt.Printf("forwarding request to address: %q\n", targetServer.Address())

	targetServer.Serve(res, req)
}
