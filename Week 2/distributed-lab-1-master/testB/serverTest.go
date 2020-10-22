package main

import(
	"bufio"
	"fmt"
	"net"
)

func readMessage(conn *net.Conn){
	reader := bufio.NewReader(*conn)
	for{
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
		fmt.Fprintln(*conn, "msg recieved")
	}
}

func main(){

	ln, _ := net.Listen("tcp", ":6969")
	for{
		conn, _ := ln.Accept()
		go readMessage(&conn)
	}
}
