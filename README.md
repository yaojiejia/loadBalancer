# Go Load Balancer

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

1. **Clone the Repository**

   ```sh
   git clone https://github.com/yaojiejia/loadBalancer.git
   cd loadBalancer
