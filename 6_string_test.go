package main

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	var dataType string = "ricid kumbara"

	fmt.Println(dataType)
	fmt.Println("String length : ", len(dataType))
	fmt.Println("Get index 0 : ", dataType[0])
	fmt.Println("Get index 0 convert to string: ", string(dataType[0]))
}
