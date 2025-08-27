package main

// listen on port 8080
// accept connection
// read request
// parse HTTP headers
// send HTTP response
// close connection

import (
	"flag"
)

func main() {

	host := flag.String("host", "127.0.0.1", "Server host")
	port := flag.Int("port", 8080, "Server port")
	flag.Parse() // parse the command-line flags

	ServerForever(*host, *port)
}
