package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var result any = random()
	// var resultString = result.(string)
	// fmt.Println(resultString)

	// var resultInt = result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case int:
		fmt.Println("Int", value)
	case string:
		fmt.Println("String", value)
	default:
		fmt.Println("Unknown", value)
	}
}
