package main

import (
	"fmt"
	
)

func main(){
	block := make(chan string)
	for i:=0 ; i<5 ; i++{
		go func(num int){
			block <- fmt.Sprintf("%s %d", "Hello world from process", num) 
		}(i)
	}

	for i:=0; i<5; i++{
		fmt.Printf("%q\n", <-block)
	}
}