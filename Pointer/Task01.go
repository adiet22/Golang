package main

import (
	"fmt"
	"math"
)

func Pembulatan(angka float64) {
	res2 := math.Round(angka*10) / 10

	fmt.Printf("%g0\n", res2)
}
