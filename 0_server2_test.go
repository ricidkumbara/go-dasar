package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer2(t *testing.T) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello world")
	})

	http.ListenAndServe(":8080", nil)
}
