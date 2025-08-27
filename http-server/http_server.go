package main

import (
	"fmt"
	"net"
)

func ServerForever(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	listner, error := net.Listen("tcp", address)
	if error != nil {
		fmt.Print(error)
	}
	for {
		conn, error := listner.Accept()
		if error != nil {
			fmt.Print(error)
		}
		buffer := make([]byte, 4096) // 1kb buffer

		n, error := conn.Read(buffer)
		if error != nil {
			// conn.Close()
		}
		if n > 0 {
			byte_read := buffer[:n]
			fmt.Print(byte_read)
			responseBytes := Handle(byte_read)
			conn.Write(responseBytes)
			// conn.Close()
		}
	}

}
