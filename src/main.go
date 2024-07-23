package main

import (
	"fmt"
	"net/http"
)

func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.go.dev"),
		newSimpleServer("https://www.spotify.com"),
	}

	// var i int
	// fmt.Println("Choose a method: round robin(1), weighted round-robin(2), IP hash(3), sticky round robin(4)")
	// fmt.Scan(&i)
	// fmt.Println("Initializing...")

	lb := newSRRLoadBalancer("8700", servers)

	handleRedirect := func(res http.ResponseWriter, req *http.Request) {
		ip := req.RemoteAddr

		lb.srrServeProxy(res, req, ip)
	}
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Program started at %s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)

}
