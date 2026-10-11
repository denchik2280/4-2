package main

import (
	"fmt"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("начали задачу %d", j)
		results <- j * j
		fmt.Printf("завершили задачу %d", j)
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}

	close(jobs)

	for r := 1; r <= 10; r++ {
		fmt.Println("результат: ", <-results)
	}
}