package main

import (
	"net/http"

	sample "github.com/yuta17/test-google-cloud-functions"
)

func main() {
	http.HandleFunc("/hello", sample.Hello)
	http.ListenAndServe(":8082", nil)
}
