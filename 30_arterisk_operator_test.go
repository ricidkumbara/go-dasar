package main

import (
	"fmt"
	"testing"
)

type Address2 struct {
	City, Provice, Country string
}

func TestArteriskOperator(t *testing.T) {
	address1 := Address2{"Purwakarta", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address2{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
