package checker

import (
	"net/http"
	"time"
	"url-checker/internal/models"
)

func CheckerURL(url string) models.CheckResult {
	timeStart := time.Now()

	resp, err := http.Get(url)
	latency := time.Since(timeStart)

	if err != nil {
		return models.CheckResult{
			URL:     url,
			Error:   err,
			Latency: latency,
		}
	}
	defer resp.Body.Close()

	return models.CheckResult{
		URL:     url,
		Status:  resp.StatusCode,
		Latency: latency,
	}
}
