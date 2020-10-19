package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

type buffer struct {
	b                 []int
	size, read, write int
}

func newBuffer(size int) buffer {
	return buffer{
		b:     make([]int, size),
		size:  size,
		read:  0,
		write: 0,
	}
}

func (buffer *buffer) get() int {
	x := buffer.b[buffer.read]
	fmt.Println("Get\t", x, "\t", buffer)
	buffer.read = (buffer.read + 1) % len(buffer.b)
	return x
}

func (buffer *buffer) put(x int) {
	buffer.b[buffer.write] = x
	fmt.Println("Put\t", x, "\t", buffer)
	buffer.write = (buffer.write + 1) % len(buffer.b)
}

func producer(buffer *buffer, start, delta int, mutex *sync.Mutex) {
	x := start
	for {
		mutex.Lock()
		buffer.put(x)
		x = x + delta
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func consumer(buffer *buffer, mutex *sync.Mutex) {
	for {
		mutex.Lock()
		_ = buffer.get()
		mutex.Unlock()
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
	}
}

func main() {
	buffer := newBuffer(5)

	var mutex = &sync.Mutex{}
	go producer(&buffer, 1, 1, mutex)
	go producer(&buffer, 1000, -1, mutex)

	consumer(&buffer, mutex)
}
