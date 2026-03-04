package service

import (
	"fmt"

	"github.com/spf13/pflag"
)

func InputHelp() {
	fmt.Println("URL Checker - утилита для параллельной проверки доступности URL")
	fmt.Println("\nИспользование:")
	fmt.Println("  url-checker [флаги] [URL...]")
	fmt.Println("\nФлаги:")
	pflag.PrintDefaults()
	fmt.Println("\nПримеры:")
	fmt.Println("  url-checker -u https://google.com,https://ya.ru")
	fmt.Println("  url-checker -t 5s https://google.com https://github.com")
	fmt.Println("  url-checker -f urls.txt")
}
