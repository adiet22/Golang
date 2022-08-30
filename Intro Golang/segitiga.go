package main

import "fmt"

func Segitiga(angka int) {
	if angka > 0 {
		rows := angka
		var i, j int
		for i = rows; i >= 1; i-- {
			for space := 1; space <= rows-i; space++ {
				fmt.Print("  ")
			}
			for j = i; j <= 2*i-1; j++ {
				fmt.Printf("* ")
			}
			for j = 0; j < i-1; j++ {
				fmt.Printf("* ")
			}
			fmt.Println("")
		}
	} else {
		fmt.Println("Angka must be greater than 0")
	}

}
