package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	var i int
	var rMethod, rURI string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			rMethod = strings.Fields(ln)[0]
			rURI  = strings.Fields(ln)[1]
			fmt.Println("***METHOD", rMethod)
			fmt.Println("***URI", rURI )
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	body += "\n"
	body += rMethod
	body += "\n"
	body += rURI
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}