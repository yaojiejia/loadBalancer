package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/yaojiejia/loadBalancer/balancer"
	"github.com/yaojiejia/loadBalancer/proxy"
	"gopkg.in/yaml.v2"
)

func main() {
	fmt.Println("Welcome to load balancer")

	// Read and parse the config.yaml file
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}

	// Use the configuration values
	var servers []proxy.Server
	serverAddresses := []string{config.ProxyServer.Server1, config.ProxyServer.Server2, config.ProxyServer.Server3}

	for _, addr := range serverAddresses {
		err := proxy.ValidateURL(addr)
		if err != nil {
			fmt.Println("Error with: ", err)
			continue
		}

		tempServer, err := proxy.NewSimpleServer(addr)
		if err != nil {
			fmt.Println("Error creating server: ", err)
			continue
		}

		if tempServer.IsAlive() {
			servers = append(servers, tempServer)
			fmt.Println("Server added!")
		} else {
			fmt.Println("Server provided is currently down: " + addr)
		}
	}

	fmt.Println("Forwarded servers added! Starting load balancer server...")

	var ipLoadBalancer *balancer.IPLoadBalancer
	var rrLoadBalancer *balancer.RRLoadBalancer
	var srrLoadBalancer *balancer.SRRLoadBalancer

	switch config.Balancer.Method {
	case "iphash":
		ipLoadBalancer = balancer.NewIPLoadBalancer(fmt.Sprintf("%d", config.LocalServer.Port), servers)
	case "roundrobin":
		rrLoadBalancer = balancer.NewRRLoadBalancer(fmt.Sprintf("%d", config.LocalServer.Port), servers)
	case "sroundrobin":
		srrLoadBalancer = balancer.NewSRRLoadBalancer(fmt.Sprintf("%d", config.LocalServer.Port), servers)
	default:
		log.Fatalf("Invalid balancer method: %s", config.Balancer.Method)
	}

	handleRedirect := func(res http.ResponseWriter, req *http.Request) {
		ip := req.RemoteAddr
		fmt.Printf("Handling request from IP: %s\n", ip)
		fmt.Printf("Balancer method: %s\n", config.Balancer.Method)

		switch config.Balancer.Method {
		case "iphash":
			ipLoadBalancer.IPServeProxy(res, req, ip)
		case "roundrobin":
			rrLoadBalancer.RrServeProxy(res, req, ip)
		case "sroundrobin":
			srrLoadBalancer.SrrServeProxy(res, req, ip)
		default:
			http.Error(res, "Invalid balancer method", http.StatusInternalServerError)
		}
	}
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Program started at %d\n", config.LocalServer.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.LocalServer.Port), nil)
}
