package main

import (
	"fmt"
	"testing"
)

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups error")
	}

	// Contoh salah
	// message := recover()
	// fmt.Println("Terjadi panic", message)
}

func TestRecover(t *testing.T) {
	runApp(true)
	fmt.Println("Ricid Kumbara")
}

// Recover: untuk melanjutkan code program pada saat terjadi panic
