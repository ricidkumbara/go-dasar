package main

import (
	"fmt"
	"testing"
)

func random() any {
	return "OK"
}

func TestTypeAssertion(t *testing.T) {
	var result any = random()
	// var resultString = result.(string)
	// fmt.Println(resultString)

	// var resultInt = result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case int:
		fmt.Println("Int", value)
	case string:
		fmt.Println("String", value)
	default:
		fmt.Println("Unknown", value)
	}
}
