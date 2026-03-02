package main

import (
	"context"
	"fmt"
	"sync"
	"time"
	"url-checker/internal/checker"
	"url-checker/internal/models"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.vk.com",
		"https://httpbin.org/delay/5", // тормоз
	}

	//timeout := 5 * time.Second
	results := make(chan models.CheckResult)

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(currentURL string) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			defer cancel()

			result := checker.CheckerURL(ctx, currentURL)
			results <- result
		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Checker URL:")
	for result := range results {
		if result.Error != nil {
			fmt.Printf(" %s - Ошибка: %v (время: %v)\n",
				result.URL, result.Error, result.Latency)
		} else {
			fmt.Printf("%s - Статус: %d (время: %v)\n",
				result.URL, result.Status, result.Latency)
		}
	}
	// todo: добавить тесты
	// todo: сделать ввод урл
	// todo: сделать ридми

}
