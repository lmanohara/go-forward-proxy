package main

import "flag"

func main() {

	mappings := proxyMappings{}

	host := flag.String("host", "127.0.0.1", "Server host")
	port := flag.Int("port", 8080, "Server port")
	flag.Var(&mappings, "map", "Comma seperated reserve proxy mappings: /path=http://backend, /auth=http://auth")
	flag.Parse() // parse the command-line flags
	ProxyForever(*host, *port, mappings)
}
