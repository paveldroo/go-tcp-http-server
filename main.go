package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	defer func() {
		err := li.Close()
		if err != nil {
			panic(err)
		}
	}()

	fmt.Println("Server was started!")

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()
	h := request(conn)
	response(conn, h)
}

func request(conn net.Conn) string {
	i := 0
	s := bufio.NewScanner(conn)
	var h string
	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)
		if i == 1 {
			h = strings.Fields(ln)[1]
			fmt.Println("***URL", h)
		}
		if ln == "" {
			fmt.Println("Headers are done!")
			break
		}
		i++
	}

	return h
}

func response(conn net.Conn, s string) {
	bh2 := `<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>Http Server</title>
				</head>
				<body>
					<h2>Hello from `

	ah2 := `!</h2>
				</body>
			</html>`

	b := bh2 + s + ah2

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(b))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, b)
}
