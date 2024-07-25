package main

import (
	"fmt"
	"net/http"
	"time"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(res http.ResponseWriter, req *http.Request)
}

func (s *simpleServer) Address() string {
	return s.addr
}

func (s *simpleServer) IsAlive() bool {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	res, err := client.Get(s.addr)
	if err != nil {
		fmt.Println("Error with: ", err)
		return false

	}
	defer res.Body.Close()

	return res.StatusCode == http.StatusOK
}
func (s *simpleServer) Serve(res http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(res, req)
}
