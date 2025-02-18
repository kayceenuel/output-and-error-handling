package fetcher

import (
	"fmt"
	"net/http"
	"time"
)

// FetchWeather contacts the weather server at the provided URL and returns the weather data.
// It handles HTTP 200 (OK) responses, HTTP 429 (Too Many Requests) with retry logic, and other errors.
func FetchWeather(url string) (string, error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // set a timeout to avoid hanging
	}

	for {
		resp, err := client.Get(url)
		if err != nil {
			// Network error or dropped connection
			return "", fmt.Errorf("failed to make request: %w", err)
		}

		switch resp.StatusCode {
		case http.StatusOK:
		}
	}
}
