package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	/* Tipe data slice */
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	fmt.Println(len(months), cap(months))

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Updated"
	// fmt.Println(slice1)

	// slice1[0] = "Mei-Lagi"
	// fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println("Slice2 :", slice2)

	var slice3 = append(slice2, "Test")
	fmt.Println(slice3)
	slice3[1] = "Bukan-Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	/* Bikin slice dari awal */
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Ricid"
	newSlice[1] = "Kumbara"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("Ini array", iniArray)
	fmt.Println("Ini slice", iniSlice)
}
