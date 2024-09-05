package main

import (
	"fmt"
	"testing"
)

func sayHello() {
	fmt.Println("Hello")
}

func sayHello2(firstName string, lastName string) {
	fmt.Println("Hello", firstName+" "+lastName)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Ricid", "Kumbara"
}

/* Named return value */
func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Ricid"
	middleName = "Dummy"
	lastName = "Kumbara"

	return firstName, middleName, lastName
}

/* Variadic Function */
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func TestFunction(t *testing.T) {
	sayHello()
	sayHello2("Ricid", "Kumbara")
	fmt.Println(getHello("Ricid Kumbara"))

	result := getHello("Fulan")
	fmt.Println(result)

	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	firstName2, _ := getFullName()
	fmt.Println(firstName2)

	firstName3, middleName3, lastName3 := getCompleteName()
	fmt.Println(firstName3, middleName3, lastName3)

	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println(total)

	numbers := []int{1, 2, 3, 4, 5}
	total = sumAll(numbers...) // Kirim pake slice dengan cara konversi ke variable argument
	fmt.Println(total)
}
