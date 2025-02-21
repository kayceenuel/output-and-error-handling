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

func TestFetchWeatherTooManyRequests(t *testing.T) {
	attempt := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempt++
		if attempt == 1 {
			// First attempt returns 429 with a Retry-After header set to 1 second
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		// Second attempt returns 200 OK
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Cloudy"))
	}))
	defer ts.Close()

	result, err := FetchWeather(ts.URL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result != "Cloudy" {
		t.Errorf("expected %q, got %q", "Cloudy", result)
	}
}

// TooManyRequest LongDelay test the case where the server returns a Retry-After header with a delay longer than 5 seconds.
func TestFetchWeatherTooManyRequestsLongDelay(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return 429 with a long delay (10 seconds) which should cause an error
		w.Header().Set("Retry-After", "10")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer ts.Close()

	_, err := FetchWeather(ts.URL)
	if err == nil {
		t.Fatal("expected an error due to long retry delay, got none")
	}
}
