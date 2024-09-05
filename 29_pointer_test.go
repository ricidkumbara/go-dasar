package main

import (
	"fmt"
	"testing"
)

type Address1 struct {
	City, Provice, Country string
}

func TestPointer(t *testing.T) {
	// Bukan pointer
	// address1 := Address{"Purwakarta", "Jawa Barat", "Indonesia"}
	// address2 := address1

	// address2.City = "Bandung"
	// fmt.Println(address1)
	// fmt.Println(address2)

	// Pointer
	address1 := Address1{"Purwakarta", "Jawa Barat", "Indonesia"}
	address2 := &address1
	var address3 *Address1 = &address1 // Bisa declare spt ini juga

	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
}
