package fetcher

import (
	"net/http"
)

type WeatherFetcher interface {
	FetcherWeather() (string, error)
}

// Main client struct for the fetcher
type Client struct {
	httpClient *http.Client
	baseURl    string
}

type RetryError struct {
	Err error
}
