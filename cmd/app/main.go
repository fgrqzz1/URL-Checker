package main

import (
	"fmt"
	"os"
	"url-checker/internal/service"
)

func main() {
	input, err := service.ParseInput()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}

	if input.ShowHelp {
		service.InputHelp()
	}

	if err := input.Validate(); err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		fmt.Fprintln(os.Stderr, "Используйте флаг -hellp для справки")
		os.Exit(1)
	}

	service.RunChecker(input)

	// todo: добавить тесты
	// todo: сделать ридми

}
