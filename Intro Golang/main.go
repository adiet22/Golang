package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	Segitiga(5)
	Segitiga(8)

	fmt.Println(genPass("a", "weaK"))
	fmt.Println(genPass("adiet95s", "strong"))

	var arr = []int{3, 1, 7, 3, 8, 9, 4, 10}
	var n int = 7
	fmt.Println(menonton(arr, n))
}
