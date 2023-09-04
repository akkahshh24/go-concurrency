package main

import (
	"fmt"
	"sync"
)

// It fills the channel with 5 "msg: i" strings and returns a receive-only channel
func boring(msg string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch) //Close the channel when the Go routine exits
		for i := 0; i <= 5; i++ {
			ch <- fmt.Sprintf("%s: %d", msg, i)
		}
	}()
	return ch
}

func main() {
	aliceCh := boring("Alice")
	bobCh := boring("Bob")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for msg := range aliceCh {
			fmt.Println(msg)
		}
	}()

	go func() {
		defer wg.Done()
		for msg := range bobCh {
			fmt.Println(msg)
		}
	}()

	wg.Wait()
	fmt.Println("Done!")
}
