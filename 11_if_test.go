package main

import (
	"fmt"
	"testing"
)

func TestIf(t *testing.T) {
	name := "fulan"

	if name == "ricid" {
		fmt.Println("Hello ricid")
	} else if name == "fulan" {
		fmt.Println("Hello fulan")
	} else {
		fmt.Println("Hello someone")
	}

	if length := len(name); length > 5 {
		fmt.Println("Panjang tidak sesuai")
	} else {
		fmt.Println("Panjang sesuai")
	}
}
