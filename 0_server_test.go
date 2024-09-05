package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	http.HandleFunc("/", homeRoute)
	http.ListenAndServe(":8080", nil)
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}
