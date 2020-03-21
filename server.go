package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var port string = "99"

func main() {

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Error to listen", err)
		os.Exit(1)
	}
	fmt.Println("listing on ", port)

	defer listener.Close()
	connInfoChan := make(chan net.Conn)
	go func() {
		var cnt int
		for conn := range connInfoChan {
			cnt++
			fmt.Printf("Accept connect cnt = %v, info: %v -> %v \n", cnt, conn.RemoteAddr(), conn.LocalAddr())
			go func(conn net.Conn) {
				defer conn.Close()
				time.Sleep(30 * time.Second)
			}(conn)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("acept fail", err)
			break
		}

		connInfoChan <- conn
	}

	fmt.Println("exit")
}
