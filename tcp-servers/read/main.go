package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8070")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Addr()
	fmt.Printf("server is listning on port %v\n", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Fprint(conn, "hello there i see you")
		conn.Close()
	}
}
