package main

import "net/http"

type wrrLoadBalancer struct {
	baseLoadBalancer
	weights map[Server]float64
}

func newWRRLoadBalancer(port string, servers []Server) *wrrLoadBalancer {
	return &wrrLoadBalancer{
		baseLoadBalancer: baseLoadBalancer{
			port:    port,
			servers: servers,
		},
		weights: make(map[Server]float64),
	}

}

func (lb *wrrLoadBalancer) getNextAvailableServer() {

}

func (lb *wrrLoadBalancer) wrrServeProxy(res http.ResponseWriter, req *http.Request) {

}
