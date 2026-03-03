package models

import "time"

type InputData struct {
	URL      []string
	Timeout  time.Duration
	FilePath string
	ShowHelp bool
}
