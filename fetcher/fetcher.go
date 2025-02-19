package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"os"
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

			// If the wait time is more than 5 seconds, give up.
			if waitDuration > 5*time.Second {
				return "", fmt.Errorf("retry delay too long (%v); giving up", waitDuration)
			}

			// Inform the user if the wait time is more than 1 secondd
			if waitDuration > 1*time.Second {
				fmt.Fprintf(os.Stderr, "Server busy. Waiting %v before retrying.../n", waitDuration)
			}

			// Wait for specified duration before retrying.
			time.Sleep(waitDuration)
			continue

		default:
			// for any unexpected http statuscodes, return an error
			resp.Body.Close()
			return "", fmt.Errorf("unexpected server response: %s", resp.Status)
		}
	}
}
