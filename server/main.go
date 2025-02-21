package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting server on :8080")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
