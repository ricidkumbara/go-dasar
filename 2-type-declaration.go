package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktp = "12312312312312"
	var married = true

	fmt.Println(ktp)
	fmt.Println(married)
}