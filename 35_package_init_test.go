package main

import (
	"fmt"
	"go-dasar/database"
	_ "go-dasar/internal" // Blank identifier, digunakan ketika import package lain tapi tidak dipakai dipackage ini langsung
	"testing"
)

func TestPackageInit(t *testing.T) {
	fmt.Println(database.GetDatabase())
}
