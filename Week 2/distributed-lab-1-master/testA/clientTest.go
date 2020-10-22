package main

import(
	"bufio"
	"fmt"
	"net"
	"os"
)

func readMessage(conn *net.Conn){
	reader := bufio.NewReader(*conn)
	msg, _ := reader.ReadString('\n')
	fmt.Printf(msg)
}

func main(){

	conn, _ := net.Dial("tcp", "127.0.0.1:6969")
	stdin := bufio.NewReader(os.Stdin)
	for{
		fmt.Println("Enter text:")
		text, _ :=stdin.ReadString('\n')
		fmt.Fprintln(conn, text)
		readMessage(&conn)
	}
}

