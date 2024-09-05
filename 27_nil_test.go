package main

import (
	"fmt"
	"testing"
)

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}

	return map[string]string{
		"name": name,
	}
}

func TestNil(t *testing.T) {
	data := NewMap("Ricid")

	if data == nil {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
