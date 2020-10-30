package main
import (
	"net/rpc"
	"flag"
//	"bufio"
//	"os"
	"secretstrings/stubs"
	"fmt"
)

func main(){
	server := flag.String("server","127.0.0.1:8030","IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	
	//Just dials to the server address so we can do funky stuff
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	request := stubs.Request{Message: "Hello"}
	response := new(stubs.Response)

	client.Call(stubs.PremiumReverseHandler, request, response)

	fmt.Println(response.Message)
}
