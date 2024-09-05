package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	names := [...]string{
		"ricid",
		"kumbara",
	}

	fmt.Println(len(names))
	fmt.Println(names)

	var product [3]string
	product[0] = "HP"
	product[1] = "HP"

	fmt.Println(product)
	fmt.Println(len(product))
}
