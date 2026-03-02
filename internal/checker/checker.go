package checker

import (
	"context"
	"net/http"
	"time"
	"url-checker/internal/models"
)

func CheckerURL(ctx context.Context, url string) models.CheckResult {
	timeStart := time.Now()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return models.CheckResult{
			URL:     url,
			Error:   err,
			Latency: time.Since(timeStart),
		}
	}

	resp, err := http.DefaultClient.Do(req)
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
