package main

import (
	"fmt"
	"testing"
)

type Customer struct {
	Name    string
	Address string
	Age     int
}

/* Stuct Method */
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, " my name is ", customer.Name)
}

func TestStruct(t *testing.T) {
	/* Cara 1 */
	customer := Customer{
		Name:    "Ricid Kumbara",
		Address: "PWK",
		Age:     25,
	}

	/* Cara 2 */
	var newCustomer Customer
	newCustomer.Name = "Fulan"
	newCustomer.Address = "KRW"
	newCustomer.Age = 28

	/* Cara 3 */
	foreignCustomer := Customer{"Fulana", "CKR", 30}

	fmt.Println(customer, newCustomer, foreignCustomer)
	customer.sayHello("Kumbara")
	newCustomer.sayHello("Kumbara")
	foreignCustomer.sayHello("Kumbara")
}

// Penggunaan struct biasa memakai PascalCase
