package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	route := request(conn)
	if route {
		fmt.Println("it was get")
	} else {
		fmt.Println("it was not get")
	}

	// write response
	respond(conn, route)
}

func request(conn net.Conn) bool {
	i := 0
	scanner := bufio.NewScanner(conn)
	var out bool
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0] // method
			u := strings.Fields(ln)[1] // uri
			out = m == "GET"
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return out
}

func respond(conn net.Conn, route bool) {
	var body string
	if route {
		body = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World with GET</strong></body></html>`
	} else {
		body = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
	}

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
	// if route {
	// 	fmt.Fprintf(conn, "it was get")
	// } else {
	// 	fmt.Fprint(conn, "it was not get")
	// }
}
