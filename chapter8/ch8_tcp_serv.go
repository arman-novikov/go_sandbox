package main
import (
	"strconv"
	"fmt"
	"net"
)

const HTML_HEADER string =
`HTTP/1.1 200 OK\r
Host: 127.0.0.1\r
Content-Type: text/html; charset=UTF-8\r
Connection: close\r
Content-Length: `

const HTML_PAGE string =
`<!DOCTYPE html>
<html>
<head>
    <title></title>
</head>
<body>
<h1>Worked!</h1>
</body>
</html>`

func server() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	page_len := strconv.FormatInt(int64(len(HTML_PAGE)), 10)
	buf := HTML_HEADER + page_len + "\r\n\r\n" + HTML_PAGE
	c.Write([]byte(buf))
	c.Close()
}

func main() {
	go server()
	var input string
	fmt.Scanln(&input)
}