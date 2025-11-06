package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/*
Building upon the code from the previous exercise:
In that previous exercise, we WROTE to the connection.

Now I want you to READ from the connection.

You can READ and WRITE to a net.Conn as a connection implements both the reader and writer interface.

Use bufio.NewScanner() to read from the connection.

After all of the reading, include these lines of code:

fmt.Println("Code got here.") io.WriteString(c, "I see you connected.")

Launch your TCP server.

In your web browser, visit localhost:8080.

Now go back and look at your terminal.

Can you answer the question as to why "I see you connected." is never written?
*/
func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	fmt.Println("server is started on port 8090")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler(conn)
	}
}
func handler(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		fmt.Fprintf(conn, "I see you are connected\n")
		fmt.Fprint(conn, "Content-Type:", "text/html")
		fmt.Println(sc.Text())
	}
	defer conn.Close()
	fmt.Print("code got here!!")
}
