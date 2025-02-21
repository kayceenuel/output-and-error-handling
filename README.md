# Output and Error Handling in Go

This project demonstrates error handling and output management in Go by building a client that interacts with a provided weather server. The client fetches weather data from the server, handles errors, and outputs messages to standard output (stdout) and standard error (stderr).

## Overview

The project consists of two main components:

- **Client**: Fetches weather data from the server, implements retry logic for rate-limited responses (HTTP 429), and properly reports errors.
- **Server**: Provided for testing, the server simulates various responses (successful weather data, HTTP 429 with a `Retry-After` header, and dropped connections).

## How to Run

### Start the Server

1. Open a terminal and navigate to the `server` folder:
   ```bash
   cd server
   ```
2. Run the server:
   ```bash
   go run main.go
   ```
   The server listens on port 8080 and simulates various responses (normal, 429, dropped connections).

### Run the Client

1. Open another terminal and navigate to the project root:
   ```bash
   cd output-and-error-handling
   ```
2. Run the client:
   ```bash
   go run main.go
   ```
   The client fetches weather data from http://localhost:8080, handles retries for 429 responses, and prints output to stdout (weather data) or stderr (error messages).

## Testing

To run the unit tests for the fetcher package:
```bash
go test ./fetcher
```