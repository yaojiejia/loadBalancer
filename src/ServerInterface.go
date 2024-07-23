package main

import "net/http"

type Server interface {
	Address() string
	IsAlive() bool
	Serve(res http.ResponseWriter, req *http.Request)
}

func (s *simpleServer) Address() string {
	return s.addr
}

func (s *simpleServer) IsAlive() bool {
	return true
}
func (s *simpleServer) Serve(res http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(res, req)
}
