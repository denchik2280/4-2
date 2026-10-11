package main

import (
	"fmt"
	"time"
)

func main() {
	sem := make(chan struct{}, 5)
	tick := time.Tick(200 * time.Millisecond)

	for i := 1; i <= 15; i++ {

		go func(id int) {
			sem <- struct{}{}
			defer func() { <-sem }()
			fmt.Println(id)
		}(i)
		<-tick
	}
}