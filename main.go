package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}
	io.WriteString(conn, "This is a Tcp Server\n")
	fmt.Fprintln(conn, "Please Write Anything To The Connection Before It Closes (10 s)")
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Fprintf(conn, "You wrote: %s\n", ln)
	}
	defer conn.Close()
}
