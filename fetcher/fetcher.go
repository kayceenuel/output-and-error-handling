package fetcher

import (
	"fmt"
	"io"
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
			// Successful response
			body, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				return "", fmt.Errorf("failed to read response body: %w", err)
			}
			return string(body), nil

		case http.StatusTooManyRequests:
			// Hanlde 429 response by reading the Retry-After header
			retryAfter := resp.Header.Get("Retry-After")
			resp.Body.Close() // won't be using the body

			// Parse the Retry-After header
			waitDuration, err := parseRetryAfter(retryAfter) // custom functionn
			if err != nil {
				// if parsing fails, default to 1 second
				fmt.Fprintln(os.stderr, "Retry-After header Invaild, waiting 1 second")
				waitDuration = 1 * time.Second
			}

		}
	}
}
