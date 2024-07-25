package balancer

// import (
// 	"net/http"

// 	"github.com/yaojiejia/loadBalancer/proxy"
// )

// type wrrLoadBalancer struct {
// 	baseLoadBalancer
// 	weights map[proxy.Server]float64
// }

// func newWRRLoadBalancer(port string, servers []proxy.Server) *wrrLoadBalancer {
// 	return &wrrLoadBalancer{
// 		baseLoadBalancer: baseLoadBalancer{
// 			port:    port,
// 			servers: servers,
// 		},
// 		weights: make(map[proxy.Server]float64),
// 	}

// }

// func (lb *wrrLoadBalancer) assignWeights() {

// }

// func (lb *wrrLoadBalancer) getNextAvailableServer() {

// }

// func (lb *wrrLoadBalancer) wrrServeProxy(res http.ResponseWriter, req *http.Request) {

// }
