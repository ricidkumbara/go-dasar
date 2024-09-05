package main

import (
	"fmt"
	"testing"
)

func endApp1() {
	fmt.Println("End app")
}

func runApp2(error bool) {
	defer endApp1()

	if error {
		panic("Ups error")
	}

	fmt.Println("App running")
}

func TestPanic(t *testing.T) {
	runApp2(true)
}

// Panic: Menghentikan program
