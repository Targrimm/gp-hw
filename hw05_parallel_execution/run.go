package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func worker(job <-chan Task, result chan<- bool, shutDown <-chan bool, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	for {
		select {
		case <-shutDown:
			return
		case task, ok := <-job:
			if !ok {
				result <- true
				continue
			}
			err := task()
			result <- err != nil
		}
	}
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var wg sync.WaitGroup
	var errorsTotal int32
	var taskCounter int32
	var resultsTotal int32
	var res error

	jobs := make(chan Task, len(tasks))
	results := make(chan bool, n)
	shutDown := make(chan bool, n)

	for range n {
		wg.Add(1)
		go worker(jobs, results, shutDown, &wg)
		if int(taskCounter) < (len(tasks) - 1) {
			jobs <- tasks[taskCounter]
			atomic.AddInt32(&taskCounter, 1)
		}
	}

loop:
	for result := range results {
		if result {
			atomic.AddInt32(&errorsTotal, 1)
		}
		atomic.AddInt32(&resultsTotal, 1)
		switch {
		case m >= 0 && int(errorsTotal) > m:
			for range n {
				shutDown <- true
			}
			res = ErrErrorsLimitExceeded
			break loop
		case int(taskCounter) < len(tasks):
			jobs <- tasks[taskCounter]
			atomic.AddInt32(&taskCounter, 1)
		case int(resultsTotal) >= len(tasks):
			for range n {
				shutDown <- true
			}
			break loop
		}
	}

	wg.Wait()
	close(jobs)
	close(results)
	close(shutDown)
	return res
}
