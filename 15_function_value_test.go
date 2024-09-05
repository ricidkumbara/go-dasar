package main

import (
	"fmt"
	"testing"
)

func getGoodBy(name string) string {
	return "Good Bye " + name
}

func TestFunctionValue(t *testing.T) {
	contoh := getGoodBy
	fmt.Println(contoh("Ricid"))
}
