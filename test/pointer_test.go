package test

import (
	"fmt"
	"testing"
)

func TestPointerBasic(t *testing.T) {
	var a int
	var b *int

	a = 1
	b = &a

	fmt.Println(&a, a)
	fmt.Println(b, *b, &b)
	fmt.Println("")

	a = 2
	fmt.Println(&a, a)
	fmt.Println(b, *b, &b)
	fmt.Println("")

	*b = 3
	fmt.Println(&a, a)
	fmt.Println(b, *b, &b)
	fmt.Println("")
}

type Person struct {
	Name string
	Age  int
}

func TestPointerStruct(t *testing.T) {
	var person Person
	var newPerson *Person

	person = Person{
		Name: "Ricid",
		Age:  25,
	}

	newPerson = &person
	fmt.Println(person, &person)
	fmt.Println(*newPerson, newPerson, &newPerson)
	fmt.Println("")

	person.Name = "Ricid 0001"
	fmt.Println(person, &person)
	fmt.Println(*newPerson, newPerson, &newPerson)
	fmt.Println("")

	newPerson.Name = "Ricid 002"
	fmt.Println(person, &person)
	fmt.Println(*newPerson, newPerson, &newPerson)
	fmt.Println("")
}
