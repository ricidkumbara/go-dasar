package main

import (
	"fmt"
	"testing"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "ricid" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func TestErrorCustom(t *testing.T) {
	err := SaveData("", nil)
	if err != nil {
		// Pakai IF
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation error:", validationErr.Error())
		// } else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("Unknown Error", err.Error())
		// }

		// Pakai Switch
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("Unknown Error", finalError.Error())
		}

	} else {
		fmt.Println("OK")
	}
}
