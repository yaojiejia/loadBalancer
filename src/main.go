package main

import (
	"fmt"
	"net/http"
)

func main() {
	var servers []Server

	for len(servers) <= 3 {
		fmt.Println("Enter the targeted ip address in https or http format")
		var tempAddr string
		fmt.Scan(&tempAddr)
		err := validateURL(tempAddr)
		if err != nil {
			fmt.Println("Error with: ", err)
			continue
		}

		tempServer, err := newSimpleServer(tempAddr)

		if err != nil {
			fmt.Println("Error creating server: ", err)
		}
		if tempServer.IsAlive() {
			servers = append(servers, tempServer)
			fmt.Println("Server added!")
		} else {
			fmt.Println("Server provided is currently down, try again later or provide another server address")
		}

	}

	fmt.Println("Forwarded servers added! Starting load balancer server...")
	// servers := []Server{
	// 	newSimpleServer("https://www.facebook.com"),
	// 	newSimpleServer("https://www.youtube.con"),
	// 	newSimpleServer("https://www.spotify.com"),
	// }

	// var i int
	// fmt.Println("Choose a method: round robin(1), weighted round-robin(2), IP hash(3), sticky round robin(4)")
	// fmt.Scan(&i)
	// fmt.Println("Initializing...")

	lb := newIPLoadBalancer("8700", servers)

	handleRedirect := func(res http.ResponseWriter, req *http.Request) {
		ip := req.RemoteAddr

		lb.IPServeProxy(res, req, ip)
	}
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Program started at %s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)

}
