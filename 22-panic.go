package main

import "fmt"

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ups error")
	}

	fmt.Println("App running")
}

func main() {
	runApp(true)
}

// Panic: Menghentikan program
