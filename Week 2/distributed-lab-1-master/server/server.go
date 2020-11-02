package main

import (
	"bufio"
	"flag"
	"net"
	"fmt"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	panic(err)
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
	for{
		conn, err := ln.Accept()
		fmt.Println("Server has got a connection!")
		if err != nil{
			handleError(err)
		}else{
			conns <- conn
		}
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	reader := bufio.NewReader(client)
	for{
		msg, _ := reader.ReadString('\n')
		fmt.Println("What is")
		msgs <- Message{sender: clientid, message: msg}
	}
}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	fmt.Println("Started running")
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, _ := net.Listen("tcp", *portPtr)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	currentClientNum := 0

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients channel
			// - start to asynchronously handle messages from this client
			clients[currentClientNum] = conn
			go handleClient(conn, currentClientNum, msgs)
			currentClientNum += 1
		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
			for i:=0; i<currentClientNum; i++{
				if i != msg.sender{
					fmt.Fprintln(clients[i], msg.message)
				}
			}
		}
	}
}
