package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond) // Simulate some work by the producer
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Consuming %d\n", num)
		time.Sleep(1 * time.Second) // Simulate some work by the consumer
	}
}

func main() {
	normalCh := make(chan int)
	bufferedCh := make(chan int, 3)

	fmt.Println("Using normal channel:")
	go producer(normalCh)
	consumer(normalCh)

	time.Sleep(2 * time.Second)

	fmt.Println("Using buffered channel:")
	go producer(bufferedCh)
	consumer(bufferedCh)
}
