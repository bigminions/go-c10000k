package main

import (
	"fmt"
	"net"
)

var host string = "127.0.0.1:99"

// Remeber set os client max limt to max
func main() {
	fmt.Println("start to connect ", host)

	for {
		_, err := net.Dial("tcp", host)
		if err != nil {
			fmt.Println("fail to conn", err)
			break
		}

	}

	fmt.Println("exit")
}
