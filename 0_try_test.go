package main

import (
	"fmt"
	"testing"
)

func TestTry(t *testing.T) {
	/* Lets try array */
	var products [10]string

	products[0] = "HP"
	products[1] = "Laptop"

	fmt.Println(products[0], products[1])

	var users = [5]string{
		"ricidkumb_1",
		"ricidkumb_2",
		"ricidkumb_3",
		"ricidkumb_4",
		"ricidkumb_5",
	}

	country := [2]string{
		"Indonesia", "Malaysia",
	}

	fmt.Println(users)
	fmt.Println(country)

	/* Lets try unary operator */
	id := 0
	id++

	fmt.Println(id)

	/* Lets try comparasion operator */
	isTrue := "user" == "user!"

	fmt.Println(isTrue)

	/* Lets try boolean operator */
	// isTrue2 := "user" == "user" && "admin" == "admin"

	// fmt.Println(isTrue2)

	middleScore := 80
	finalScore := 70
	isPass := middleScore >= 80 && finalScore >= 80

	fmt.Println("Status : ", isPass)

	/* Lets try map */
	persons := map[string]string{}
	person1s := make(map[string]string)
	var person2s map[string]string
	var person3s map[string]string = map[string]string{}

	fmt.Println(persons)
	fmt.Println(person1s)
	fmt.Println(person2s)
	fmt.Println(person3s)
}
