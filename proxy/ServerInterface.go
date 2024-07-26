package proxy

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

func (s *SimpleServer) Address() string {
	return s.Addr
}

func (s *SimpleServer) IsAlive() bool {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	res, err := client.Get(s.Addr)
	if err != nil {
		fmt.Println("Error with: ", err)
		return false

	}
	defer res.Body.Close()

	return res.StatusCode == http.StatusOK
}
func (s *SimpleServer) Serve(res http.ResponseWriter, req *http.Request) {
	s.Proxy.ServeHTTP(res, req)
}
