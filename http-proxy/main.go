package main

import "flag"

func main() {

	host := flag.String("host", "127.0.0.1", "Server host")
	port := flag.Int("port", 8080, "Server port")
	flag.Parse() // parse the command-line flags
	ProxyForever(*host, *port)
}
