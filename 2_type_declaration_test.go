package main

import (
	"fmt"
	"testing"
)

func TestTypeDeclaration(t *testing.T) {
	type NoKTP string
	type Married bool

	var ktp NoKTP = "12312312312312"
	var married Married = true

	fmt.Println(ktp)
	fmt.Println(married)
}
