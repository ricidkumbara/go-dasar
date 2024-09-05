package main

import (
	"errors"
	"fmt"
	"testing"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	}

	return nilai / pembagi, nil
}

func TestError(t *testing.T) {
	hasil, err := Pembagian(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println(err)
	}
}
