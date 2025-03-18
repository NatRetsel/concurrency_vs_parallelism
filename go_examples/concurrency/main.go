package main

import (
	"fmt"
	"time"
)

func sayHello(goroutineNum string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%v: Hello from goroutine! - count: %v\n", goroutineNum, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	fmt.Println("==========Single Print go routine example==============")

	go sayHello("goroutine 1") // Launches sayHello in a goroutine
	go sayHello("goroutine 2")
	fmt.Println("Hello from main!")
	time.Sleep(time.Second * 2)
	fmt.Println("=======================================================")

	fmt.Println("==========Print from channel go routine example==============")
	ch := make(chan string)
	go func() {
		ch <- "Hello from goroutine channel!"
		close(ch)
	}()

	msg, ok := <-ch
	if !ok {
		fmt.Println("Channel closed")
	} else {
		fmt.Printf("%v\n", msg)
	}

	fmt.Println("=============================================================")
}
