package main

import (
	"fmt"
	"go-dasar/helper"
	"testing"
)

func TestImport(t *testing.T) {
	result := helper.SayHello("Ricid")
	fmt.Print(result)
}
