package main

import (
	"fmt"
	"testing"
)

func TestConstant(t *testing.T) {
	const name string = "ricid"
	const (
		firstName = "Ricid"
		lastName  = "Kumbara"
	)

	fmt.Println(name, firstName, lastName)
}
