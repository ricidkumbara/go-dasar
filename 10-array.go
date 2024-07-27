package main

import "fmt"

func main() {
	names := [...]string{
		"ricid",
		"kumbara",
	}

	fmt.Println(len(names))

	var product [3]string
	product[0] = "HP"
	product[1] = "HP"

	fmt.Println(product)
	fmt.Println(len(product))
}
