package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApp() {
	defer logging()
	fmt.Println("App running")
}

func main() {
	runApp()
}

// Fitur untuk memanggil function lain ketika function saat ini selesai
// Sekalipun aplikasi error, defer akan tetap dieksekusi
