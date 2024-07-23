package main

import "net/http"

type Server interface {
	Address() string
	IsAlive() bool
	Serve(res http.ResponseWriter, req *http.Request)
}
