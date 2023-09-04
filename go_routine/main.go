package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for ch := 'a'; ch <= 'e'; ch++ {
		fmt.Printf("%c\n", ch)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go printNumbers() // start a new Go routine for printNumbers()
	printLetters()    // Execute printLetters() concurrently with the Go routine
	fmt.Print("Done!")
}
