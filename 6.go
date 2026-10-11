package main

import (
	"fmt"
	"time"
)

func main() {
	sources := []string{"база1", "база2", "база3", "база4"}
	responses := make(chan string, len(sources))

	for _, b := range sources {
		go func(s string) {
			if s == "база2" {
				time.Sleep(100 * time.Millisecond)
			} else {
				time.Sleep(300 * time.Millisecond)
			}

			responses <- "результат из " + s
		}(b)
	}

	fastest := <-responses
	fmt.Println("первый успешный результат:", fastest)
}