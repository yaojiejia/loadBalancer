# Go Load Balancer
![Static Badge](https://img.shields.io/badge/go%20version-1.20-green)
[![CodeFactor](https://www.codefactor.io/repository/github/yaojiejia/loadbalancer/badge)](https://www.codefactor.io/repository/github/yaojiejia/loadbalancer)
[![Go Report Card](https://goreportcard.com/badge/github.com/yaojiejia/loadBalancer)](https://goreportcard.com/report/github.com/yaojiejia/loadBalancer)
![Static Badge](https://img.shields.io/badge/license-MIT-green)\
A lightweight and efficient load balancer written in Go, supporting multiple load balancing algorithms including Round Robin, Weighted Round Robin, Sticky Round Robin, and IP Hashing.

## Features

- **Round Robin**: Distributes requests evenly across all servers.
- **Weighted Round Robin**: Distributes requests based on server weights, allowing for more requests to be sent to more powerful servers.
- **Sticky Round Robin**: Maintains session persistence, sending requests from the same client to the same server.
- **IP Hashing**: Uses the client's IP address to determine which server will handle the request, ensuring that requests from the same IP address are consistently sent to the same server.

## Setup

### Prerequisites

- [Docker](https://www.docker.com/products/docker-desktop) installed on your machine.

### Files

- **main.go**: The main Go application file.
- **Dockerfile**: The Dockerfile to build the Docker image.
- **go.mod**: The Go module file.
- **loadBalancer**: Directory containing configuration for load balancing

### Building and Running the Load Balancer

1. **Clone the Repository and Run Locally**

   ```sh
   git clone https://github.com/yaojiejia/loadBalancer.git
   cd loadBalancer
   go run .
2. **Build the Docker Image and Run it**
   ```sh
   docker build -t loadbalancer:latest .
   docker run -d -p 8070:8070 loadbalancer:latest
   
## Configuration

### main.go

The main Go application file contains the implementation of the load balancing algorithms. You can configure the servers and the load balancing algorithm by modifying this file.

```go
package main

import (
    "fmt"
    "net/http"
    // other necessary imports
)

func main() {
    // setup your load balancing algorithms and servers here

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // load balancing logic
    })

    fmt.Println("Starting load balancer on port 8080")
    http.ListenAndServe(":8080", nil)
}
```
### config.yaml

You can also directly modify `config.yaml` to change the target IP address as well as the load balancing methods
```yaml
localServer:
  port: 8080
  host: localhost

ProxyServer:
  server1: "https://facebook.com"
  server2: "https://google.com"
  server3: "https://instagram.com"
balancer:
  method: round-robin
```
