package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:65533")
	if err != nil {
		log.Fatalf("listen : %v\n", err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept : %v\n", err)
			continue
		}
		ch := make(chan string, 10)
		go recvConn(conn, ch)
		go sendConn(conn, ch)
	}
}

func recvConn(conn net.Conn, ch chan<- string) {
	defer conn.Close()
	defer close(ch)

	rd := bufio.NewReader(conn)

	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			log.Printf("read: %v\n", err)
			return
		}

		ch <- string(line)
	}
}

func sendConn(conn net.Conn, ch <- chan string) {
	wr := bufio.NewWriter(conn)

	for {
		msg, ok := <- ch
		if !ok {
			break
		}
		wr.WriteString(msg)
		wr.Flush()
	}
}
