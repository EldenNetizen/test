package main

import (
	"fmt"
	"time"
)

func threadTest(ch chan int) {
	for i := 0; i < 5; i++ {
		fmt.Println("Thread Test: ", i)
		ch <- i
		time.Sleep(5 * time.Second)
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go threadTest(ch)
	for v := range ch {
		fmt.Println("Received from channel: ", v)
	}
}
