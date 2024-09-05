package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	person := map[string]string{
		"name":    "Ricid",
		"address": "Purwakarta",
	}

	person["title"] = "Staff IT"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Ricid"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
