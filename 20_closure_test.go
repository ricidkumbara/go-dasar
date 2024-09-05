package main

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {
	counter := 0

	increment := func() {
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
