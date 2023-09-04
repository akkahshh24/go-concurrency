package main

import (
	"fmt"
	"sync"
	"time"
)

func source(tasks chan<- int, numTasks int) {
	for i := 1; i <= numTasks; i++ {
		tasks <- i
	}
	close(tasks)
}

func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Worker %d processing task %d\n", id, task)
		time.Sleep(time.Second * time.Duration(task))
		results <- task * 2
	}
}

func main() {
	const numWorkers = 3
	const numTasks = 10

	taskQueue := make(chan int, numTasks)
	resultQueue := make(chan int, numTasks)
	var wg sync.WaitGroup

	// start the source go routine to generate tasks
	go source(taskQueue, numTasks)

	// create worker go routines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, taskQueue, resultQueue, &wg)
	}

	// wait for all worker go routines to finish
	go func() {
		wg.Wait()
		close(resultQueue)
	}()

	// print the result from the resultQueue
	for result := range resultQueue {
		fmt.Printf("Result Received: %d\n", result)
	}

	fmt.Println("All tasks done!")
}
