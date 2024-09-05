package main

import (
	"fmt"
	"testing"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func TestAnonimousFunction(t *testing.T) {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("ricid", blacklist)

	// cara penggunaan yang lain
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})
}
