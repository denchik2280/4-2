package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://discord.com",
		"https://yandex.ru",
		"https://kentstudios.ru",
	}
	jobs := make(chan string, len(urls))
	for _, u := range urls {
		jobs <- u
	}
	close(jobs)

	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for url := range jobs {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Воркер %d: %s Ошибка\n", w, url)
					continue
				}
				fmt.Printf("Воркер %d: %s %s\n", w, url, resp.Status)
				resp.Body.Close()
			}
		}()
	}
	wg.Wait()
	fmt.Println("конец")
}