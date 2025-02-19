package main

import (
	"fmt"
	"os"

	"output-and-error-handling/fetcher"
)

func main() {
	url := "http://localhost:8080" // server to run on this URL

	weather, err := fetcher.FetchWeather(url)
	if err != nil {
		// Print error to stderr and exit withh nonzero codee
		fmt.Fprintln(os.Stderr, "Error fetching weather:", err)
		os.Exit(1)
	}

	// print the weather data to stdout
	fmt.Println(weather)
}
