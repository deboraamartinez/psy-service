package main

import (
	"net/http"
)

func main() {
	// Run the server
	http.ListenAndServe(":8080", nil)
}
