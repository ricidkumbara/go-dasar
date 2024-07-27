package main

import "fmt"

func getGoodBy(name string) string {
	return "Good Bye " + name
}

func main() {
	contoh := getGoodBy
	fmt.Println(contoh("Ricid"))
}
