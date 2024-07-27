package main

import "fmt"

func main() {
	var dataType string

	dataType = "ricid kumbara"

	fmt.Println(dataType)
	fmt.Println("String length : ", len(dataType))
	fmt.Println("Get index 0 : ", dataType[0])
	fmt.Println("Get index 0 convert to string: ", string(dataType[0]))
}
