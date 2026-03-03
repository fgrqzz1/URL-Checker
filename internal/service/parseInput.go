package service

import (
	"fmt"
	"github.com/spf13/pflag"
	"time"
	"url-checker/internal/models"
)

func ParseInput() (*models.InputData, error) {
	input := models.InputData{
		Timeout: time.Second * 5,
	}

	pflag.StringSliceVarP(&input.URL, "url", "u", []string{}, "URLs для проверки (через запятую)")
	pflag.DurationVarP(&input.Timeout, "timeout", "t", input.Timeout, "таймаут на запрос")
	pflag.StringVarP(&input.FilePath, "file", "f", "", "файл со списком URL (по одному на строку)")
	pflag.BoolVarP(&input.ShowHelp, "help", "h", false, "показать помощь")

	pflag.Parse()

	input.URL = append(input.URL, pflag.Args()...)

	if input.FilePath != "" {
		fileURLs, err := readURLsFromFile(input.FilePath)
		if err != nil {
			return nil, fmt.Errorf("Ошибка чтения файла: %v", err)
		}
		input.URL = append(input.URL, fileURLs...)
		fmt.Printf("Добавлено %d URL из файла\n", len(fileURLs))
	}

	if len(input.URL) == 0 {
		return nil, fmt.Errorf("Не указаны URL для проверки")
	}

	return &input, nil
}
