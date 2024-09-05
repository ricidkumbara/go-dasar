package main

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T) {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	names := []string{"ricid", "kumbara"}
	for j := 0; j < len(names); j++ {
		fmt.Println(names[j])
	}

	for index, name := range names {
		fmt.Println("index", index, " = ", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke ", i)
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println("Perulangan ke ", i)
	}

}
