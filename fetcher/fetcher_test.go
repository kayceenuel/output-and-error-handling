package fetcher

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFetchWeatherOK tests that a 200 OK response returns the expeted weather data.
func TestFetchWeatherOK(t *testing.T) {
	expected := "Sunny and 75°F"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(expected))
	}))
	defer ts.Close()

	result, err := FetchWeather(ts.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

// TestFetchWeatherTooManyRequests tests the retry mechanism when a 429 response is received.

func TestFetchWeatherTooManyRequest(t *testing.T) {
	attempt := 0
	ts := http.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempt++
		if attempt == 1 {
			// first attempt returns 429 with a RetryAfter header set to 1 second
			w.Header().Set("Retry-After", "1")
			w.WriterHeader(http.StatusTooManyRequests)
			return
		}
	}))
}
