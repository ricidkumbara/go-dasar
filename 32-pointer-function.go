package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// address := Address{}
	// changeAddressToIndonesia(&address)

	var address *Address = &Address{}
	changeAddressToIndonesia(address)

	fmt.Println(address)
}
