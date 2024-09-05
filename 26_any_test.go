package main

import (
	"fmt"
	"testing"
)

func Ups() any {
	//return 1
	//return true
	return "ups"
}

func TestAny(t *testing.T) {
	var kosong = Ups()
	fmt.Println(kosong)
}
