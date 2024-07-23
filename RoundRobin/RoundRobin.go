package main

type rbLoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}
