package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8083")
	if err != nil {
		fmt.Println("error starting server:", err.Error())
	}
	conn, _ := ln.Accept()

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			return
		}

		fmt.Println("Server: ", msg)
		conn.Write([]byte(strings.ToUpper(msg)))
	}
}
