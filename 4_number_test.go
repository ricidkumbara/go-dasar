package main

import (
	"fmt"
	"testing"
)

func TestNumber(t *testing.T) {
	// punya alias byte
	var dataTypeInt8 int8

	// punya alias rune
	var dataTypeInt16 int16

	var dataTypeInt32 int32
	var dataTypeInt64 int32

	dataTypeInt8 = 1
	dataTypeInt16 = 1
	dataTypeInt32 = 1
	dataTypeInt64 = 1

	fmt.Println("Output : ", dataTypeInt8)
	fmt.Println("Output : ", dataTypeInt16)
	fmt.Println("Output : ", dataTypeInt32)
	fmt.Println("Output : ", dataTypeInt64)
}
