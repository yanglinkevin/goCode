package main

import (
	"fmt"
	"net"
	"os"
)

func main () {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("dial failed:", err)
		os.Exit(1)
	}
	for {
	var buffer []byte = []byte("You are welcome. I'm server.")
	n, err2 := conn.Write(buffer)
	if err2 != nil {
		fmt.Println("Read failed:", err2)
		return
	}
	fmt.Println("Read success:", n)
	//defer conn.Close()
	}
}
