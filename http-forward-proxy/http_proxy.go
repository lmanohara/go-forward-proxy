package main

import (
	"fmt"
	"net"
)

func ProxyForever(host string, port int) {

	address := fmt.Sprintf("%s:%d", host, port)
	listener, error := net.Listen("tcp", address)
	if error != nil {
		fmt.Println(error)
	}

	for {
		conn, error := listener.Accept()
		if error != nil {
			fmt.Println(error)
		}
		buffer := make([]byte, 4096)
		n, error := conn.Read(buffer)
		if error != nil {
			fmt.Println(error)
		}

		if n > 0 {
			byte_read := buffer[:n]
			out := Handle(byte_read)
			fmt.Print(byte_read)
			conn.Write(out)
			// // conn.Close()
		}
	}

}
