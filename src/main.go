package main

import (
	"fmt"
	"handler"
	"net"
	"os"
	"protocol"
)

func main() {
	netListen, err := net.Listen("tcp", "localhost:9009")
	if err != nil {
		os.Exit(1)
	}

	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	tmpBuffer := make([]byte, 0)
	readerChannel := make(chan []byte, 16)
	go reader(readerChannel)
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		tmpBuffer = protocol.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}

func reader(readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			handler.HandleData(data)
		}
	}
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}
