package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

/*
Create a basic server using TCP.
The server should use net.Listen to listen on port 8080.

Remember to close the listener using defer.

Remember that from the "net" package you first need to LISTEN, then you need to ACCEPT an incoming connection.

Now write a response back on the connection.

Use io.WriteString to write the response: I see you connected.

Remember to close the connection.

Once you have all of that working, run your TCP server and test it from telnet (telnet localhost 8080).
*/
func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		t := 10
		ticker := time.NewTicker(1 * time.Second)
		fmt.Fprint(conn, "I see You Are connected\n")
		io.WriteString(conn, "hello world")
		for i := t; i >= 0; i-- {
			fmt.Fprintf(conn, "disconnecting in : %v\n", i)
			<-ticker.C

		}
		conn.Close()
	}

}
