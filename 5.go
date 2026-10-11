package main

import (
	"crypto/md5"
	"fmt"
	"sync"
	"os"
)

func main() {
	files := []string{"mama.txt", "papa.txt", "buba.txt", "yooy.txt"}
	jobs := make(chan string, len(files))
	for _, f := range files {
		jobs <- f
	}
	close(jobs) 

var wg sync.WaitGroup


	for w := 1; w <= 2; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for filename := range jobs {
			
				data, err := os.ReadFile(filename)
				if err != nil {
					fmt.Printf("Воркер %d: ошибка %s (%v)\n", w, filename, err)
					continue
				}

				hash := md5.Sum(data)
				fmt.Printf("Воркер %d: %s %x\n", w, filename, hash)
			}
		}()
	}

	wg.Wait()
	fmt.Println("Все!")
}