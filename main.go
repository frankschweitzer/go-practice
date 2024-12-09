package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := ":8080"
	fmt.Printf("Starting server on port %s \n", port)
	http.HandleFunc("/", handler)
	http.ListenAndServe(port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "FIRST GO WEBSITE")
}
