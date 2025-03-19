package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(goroutineNum string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v: Hello from goroutine! - count: %v\n", goroutineNum, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func sendData(ch chan<- string, data string) {
	fmt.Println("Sending data: ", data)
	ch <- data // blocks until data is sent
	fmt.Println("Finished sending: ", data)
}

func receiveData(ch <-chan string) {
	data := <-ch // blocks until data is received
	fmt.Println("Received: ", data)
}

type sharedObj struct {
	mu    sync.Mutex
	count int
}

func (so *sharedObj) increment() {
	so.mu.Lock()
	so.count++
	fmt.Printf("Count: %v\n", so.count)
	so.mu.Unlock()
}

func main() {
	fmt.Println("==========Single Print go routine example==============")
	go sayHello("goroutine 1") // Launches sayHello in a goroutine
	go sayHello("goroutine 2")
	fmt.Println("Hello from main!")
	time.Sleep(time.Second * 2)
	fmt.Println("=======================================================")

	fmt.Println("=======Unbuffered Channel go routine example===========")
	ch := make(chan string) // Unbuffered channel
	go sendData(ch, "dummy data")
	go receiveData(ch)
	time.Sleep(time.Second)
	fmt.Println("=======================================================")

	fmt.Println("===================Mutex example=======================")
	so := sharedObj{
		count: 0, // count initialized to zero value
	}

	// anonymous function

	for i := 0; i < 1000; i++ {
		go func() {
			so.increment()
			time.Sleep(time.Millisecond * 100)
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("Final count: %v\n", so.count)
	fmt.Println("=======================================================")
}
