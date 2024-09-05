package main

import (
	"fmt"
	"testing"
)

func sayHelloWithFilter1(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter1(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}

func TestFunctionParam(t *testing.T) {
	filter := spamFilter1
	sayHelloWithFilter1("Anjing", filter)
}
