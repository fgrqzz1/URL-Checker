package main

import (
	"context"
	"fmt"
	"time"
	"url-checker/internal/checker"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.vk.com",
		"https://httpbin.org/delay/5", // тормоз
	}

	for _, url := range urls {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

		result := checker.CheckerURL(ctx, url)
		cancel()

		fmt.Println("Checker URL:")
		if result.Error != nil {
			fmt.Printf(" %s - Ошибка: %v (время: %v)\n",
				result.URL, result.Error, result.Latency)
		} else {
			fmt.Printf("%s - Статус: %d (время: %v)\n",
				result.URL, result.Status, result.Latency)
		}
	}
}
