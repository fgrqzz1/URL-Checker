package main

import (
	"context"
	"fmt"
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

	for _, url := range urls {
		go func(currentURL string) {
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

			result := checker.CheckerURL(ctx, currentURL)
			cancel()
			results <- result
		}(url)
	}

	fmt.Println("Checker URL:")
	for i := 0; i < len(urls); i++ {
		result := <-results
		close(results)

		if result.Error != nil {
			fmt.Printf(" %s - Ошибка: %v (время: %v)\n",
				result.URL, result.Error, result.Latency)
		} else {
			fmt.Printf("%s - Статус: %d (время: %v)\n",
				result.URL, result.Status, result.Latency)
		}
	}

}
