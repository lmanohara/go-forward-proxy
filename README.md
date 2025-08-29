# HTTP Proxy Suite

**HTTPProxySuite** is a Go project demonstrating a full HTTP request flow using custom proxies and servers built with standard libraries.  

The suite includes:

1. **Forward Proxy** – receives client requests and forwards them to the reverse proxy.  
2. **Reverse Proxy** – receives requests from the forward proxy and forwards them to the HTTP server.  
3. **HTTP Server** – processes requests and returns responses back through the proxy chain.  

This project showcases lightweight, end-to-end request handling, proxying, and server communication in Go, all without third-party frameworks.

---

## Project Structure


- `http-forward-proxy/` – contains the forward proxy server code.  
- `http-reverse-proxy/` – contains the reverse proxy server code.  
- `http-server/` – contains the simple HTTP server code and index files for each server.  
- `docker-compose.yml` – orchestrates the containers and networks.

---

## Requirements

- [Docker](https://www.docker.com/get-started) >= 20.x  
- [Docker Compose](https://docs.docker.com/compose/) >= 1.29.x  
- Go >= 1.20 (for building the binaries)

---

## Setup & Running

### 1. Build and start all services