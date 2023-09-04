package main

import (
	"fmt"
	"sync"
	"time"
)

// job represents the task to be executed by a worker
type job struct {
	ID int
}

// workerPool represents a pool of worker Go routines
type workerPool struct {
	numWorkers int
	jobQueue   chan job
	results    chan int
	wg         sync.WaitGroup
}

// newWorkerPool creates a new worker pool with the specified number of workers
func newWorkerPool(numWorkers, jobQueueSize int) *workerPool {
	return &workerPool{
		numWorkers: numWorkers,
		jobQueue:   make(chan job, jobQueueSize),
		results:    make(chan int, jobQueueSize),
	}
}

// addJob adds a job to the job queue
func (wp *workerPool) addJob(job job) {
	wp.jobQueue <- job
}

// start starts the worker pool and dispatches jobs to workers
func (wp *workerPool) start() {
	for i := 1; i <= wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker function to process jobs from the queue
func (wp *workerPool) worker(id int) {
	defer wp.wg.Done()
	for job := range wp.jobQueue {
		fmt.Printf("Worker %d started job %d\n", id, job.ID)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
		wp.results <- job.ID
	}
}

// Wait waits for all worker go routines to finish and closes the results channel
func (wp *workerPool) wait() {
	wp.wg.Wait()
	close(wp.results)
}

// collectResults collects and prints the results from the results channel
func (wp *workerPool) collectResults() {
	for result := range wp.results {
		fmt.Printf("Result received for job %d\n", result)
	}
}

func main() {
	numWorkers := 3
	numJobs := 10

	workerPool := newWorkerPool(numWorkers, numJobs)

	// Adding jobs to the job queue
	for i := 1; i <= numJobs; i++ {
		workerPool.addJob(job{
			ID: i,
		})
	}

	close(workerPool.jobQueue)

	workerPool.start()
	workerPool.wait()
	workerPool.collectResults()

	fmt.Printf("Done!")
}

/*
Worker 1		Worker 2	   Worker 3
-------------------------------Job 01 start
Job 02 start-------------------------------
----------------Job 03 start---------------
-------------------------------Job 01 done
-------------------------------Job 04 start
Job 02 done--------------------------------
Job 05 start-------------------------------
----------------Job 03 done----------------
----------------Job 06 start---------------
----------------Job 06 done----------------
----------------Job 07 start---------------
-------------------------------Job 04 done
-------------------------------Job 08 start
Job 05 done--------------------------------
Job 09 start-------------------------------
-------------------------------Job 08 done
-------------------------------Job 10 start
Job 09 done--------------------------------
----------------Job 07 done----------------
--------------------------------Job 10 done
*/
