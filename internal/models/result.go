package models

import "time"

type CheckResult struct {
	URL     string        `json:"url"`
	Status  int           `json:"status,omitempty"`
	Error   error         `json:"-"`
	Latency time.Duration `json:"latency"`
}

func (r CheckResult) ErrorMessage() string {
	if r.Error != nil {
		return r.Error.Error()
	}
	return ""
}
