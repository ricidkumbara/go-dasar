package main

import "fmt"

func main() {
	name := "Ricid Kumbara"
	fmt.Println(name)

	// Test variable kosong
	var age int8
	var address string
	fmt.Println(age, address)

	person := `
		Ricid
		Kumbara
	`
	fmt.Println(person)
}
