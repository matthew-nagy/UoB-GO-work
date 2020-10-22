package main

import (
	"flag"
	"net"
	"bufio"
	"os"
	"fmt"
)

func read(conn *net.Conn) {
	reader := bufio.NewReader(*conn)
	for{
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
	}
}

func write(conn *net.Conn) {
	stdin := bufio.NewReader(os.Stdin)
	for{
		fmt.Println("Enter text:")
		text, _ := stdin.ReadString('\n')
		fmt.Fprintln(*conn, text)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	connection, err := net.Dial("tcp", *addrPtr)
	if err != nil{
		fmt.Println("Failed to connect to server")
		return
	}
	//TODO Start asynchronously reading and displaying messages
	go read(&connection)
	//TODO Start getting and sending user messages.
	write(&connection)
}
