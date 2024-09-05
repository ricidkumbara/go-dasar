package main

import (
	"fmt"
	"testing"
)

type Address struct {
	City, Provice, Country string
}

func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func TestPointerFunction(t *testing.T) {
	// address := Address{}
	// changeAddressToIndonesia(&address)

	var address *Address = &Address{}
	changeAddressToIndonesia(address)

	fmt.Println(address)
}
