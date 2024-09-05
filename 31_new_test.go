package main

import (
	"fmt"
	"testing"
)

type Address3 struct {
	City, Provice, Country string
}

func TestNew(t *testing.T) {
	alamat1 := new(Address3)
	alamat2 := alamat1

	alamat1.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
