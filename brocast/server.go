package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	s, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := s.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		upper := strings.ToUpper(scan.Text()) + "\n"
		io.WriteString(conn, upper)
	}
}
