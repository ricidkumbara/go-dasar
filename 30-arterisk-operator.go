package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func main() {
	address1 := Address{"Purwakarta", "Jawa Barat", "Indonesia"}
	address2 := &address1
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}
