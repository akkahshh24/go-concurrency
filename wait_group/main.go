package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this Go routine is done when the function returns
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done() // Notify the WaitGroup that this Go routine is done when the function returns
	for ch := 'a'; ch <= 'e'; ch++ {
		fmt.Printf("%c\n", ch)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Add the number of Go routines to wait for

	go printNumbers(&wg) // Start a new Go routine for printNumbers()
	go printLetters(&wg) // Start a new Go routine for printLetters()

	wg.Wait() // Wait until all Go routines in the WaitGroup finish
	fmt.Print("Done!")
}
