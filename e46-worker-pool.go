package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "-------------------finished job", j)
		results <- fmt.Sprintf("Job Done - %d", j)
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan string, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	for a := 1; a <= numJobs; a++ {
		fmt.Println("Output <<<< ", <-results)
	}
}

//https://go.dev/play/p/bhFfmxGK-I6
