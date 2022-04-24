package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	go recvMessage(conn)
	sendMessage(conn)
}

func recvMessage(conn net.Conn) {
	scan := bufio.NewScanner(conn)
	for scan.Scan() {

		fmt.Println(scan.Text())
	}
}

func sendMessage(conn net.Conn) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("start to scan")
	for input.Scan() {
		str := input.Text() + "\n"
		io.WriteString(conn, str)
	}
}
