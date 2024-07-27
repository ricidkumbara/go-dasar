package main

import "fmt"

type Address struct {
	City, Provice, Country string
}

func main() {
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat1.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
