package main

import (
	"fmt"
	"testing"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}

func TestFunctionTypeDeclaration(t *testing.T) {
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
