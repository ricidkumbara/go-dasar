package main

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {
	name := "fulan"

	switch name {
	case "ricid":
		fmt.Println("Hello Ricid")
	case "fulan":
		fmt.Println("Hello Fulan")
	default:
		fmt.Println("Hello Other")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	default:
		fmt.Println("Nama OK")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama telalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("nama sudah ok")
	}
}
