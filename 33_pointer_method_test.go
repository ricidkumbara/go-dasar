package main

import (
	"fmt"
	"testing"
)

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func TestPointerMethod(t *testing.T) {
	man := Man{"Fulan"}
	man.Married()

	fmt.Println(man.Name)
}
