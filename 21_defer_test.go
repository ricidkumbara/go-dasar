package main

import (
	"fmt"
	"testing"
)

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp1() {
	defer logging()
	fmt.Println("App running")
}

func TestDefer(t *testing.T) {
	runApp1()
}

// Fitur untuk memanggil function lain ketika function saat ini selesai
// Sekalipun aplikasi error, defer akan tetap dieksekusi
