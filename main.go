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
	u := request(conn)
	switch u {
	case "/":
		response(conn)
	case "/apply":
		responseApply(conn)
	default:
		response(conn)
	}
}

func request(conn net.Conn) string {
	i := 0
	s := bufio.NewScanner(conn)
	var u string
	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)
		if i == 0 {
			u = strings.Fields(ln)[1]
			fmt.Println("***URI", u)
		}
		if ln == "" {
			fmt.Println("Headers are done!")
			break
		}
		i++
	}

	return u
}

func response(conn net.Conn) {
	b := `<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>Http Server</title>
				</head>
				<body>
					<h2>It's just a starting page!'</h2>
				</body>
			</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(b))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, b)
}

func responseApply(conn net.Conn) {
	b := `<!DOCTYPE html>
			<html lang="en">
				<head>
					<meta charset="UTF-8">
					<title>Http Server</title>
				</head>
				<body>
					<h2>You just applied this training!</h2>
				</body>
			</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(b))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, b)
}
