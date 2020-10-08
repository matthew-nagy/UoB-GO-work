package main

import (
	"log"
	"os"
	"runtime/trace"
	"time"
	"fmt"
)

func foo(channel chan string) {
	fmt.Println("foo: Sending ping")
	channel <- "ping"
	val := <- channel
	for ;val != "pong";{
		channel <- "ping"
	}

	fmt.Println("foo: Recieved pong")

	go foo(channel)
}

func bar(channel chan string) {
	val := <- channel
	fmt.Println("bar: Recieved ", val)
	fmt.Println("bar: Sending pong")
	channel <- "pong"

	go bar(channel)
}

func pingPong() {
	// TODO: make channel of type string and pass it to foo and bar
	c := make(chan string)
	go foo(c) // Nil is similar to null. Sending or receiving from a nil chan blocks forever.
	go bar(c)
	time.Sleep(500 * time.Millisecond)
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	pingPong()
}
