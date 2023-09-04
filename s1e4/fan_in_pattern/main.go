package main

import (
	"fmt"
	"sync"
)

// producer creates a channel, fills it with some values and returns a receive-only channel
func producer(id int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- id*10 + i
		}
	}()
	return ch
}

// fanIn merges the producer channels into one channel called the output channel
func fanIn(inputs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	output := make(chan int) // output channel to merge both channel values

	copy := func(ch <-chan int) { // copy function to copy the value from the producer channel to the output channel
		defer wg.Done()
		for val := range ch {
			output <- val
		}
	}

	wg.Add(len(inputs)) // add the number of channels to the waitgroup
	for _, ch := range inputs {
		go copy(ch) // spwan a copy go routine
	}

	go func() { // wait for the go routines to finish copying to the output channel and close the channel
		wg.Wait()
		close(output)
	}()

	return output
}

func main() {
	// Step-1 Producer
	ch1 := producer(1)
	ch2 := producer(20)

	// Step-2 FanIn
	mergedOutputChannel := fanIn(ch1, ch2)

	// Step-3 Consumer
	for val := range mergedOutputChannel {
		fmt.Println(val)
	}

	fmt.Println("Done!")
}
