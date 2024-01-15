package main

import (
	"fmt"
	"net"
)

func requestHandler(c net.Conn) {
	data := make([]byte, 4096)

	for {
		n, err := c.Read(data)
		if err != nil {
			return
		}

		fmt.Println(string(data[:n]))

		c.Write(data[:n])
		if err != nil {
			return
		}
	}
}

func main() {
	// 포트만 적으면 모든 IP에서 연결받음
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}
}
