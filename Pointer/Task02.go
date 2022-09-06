package main

import "fmt"

type Deret struct {
	limit int
}

func DeretBilangan(angka int) {
	var data Deret
	data.limit = angka
	fmt.Println("\nLimit n			:", angka)
	fmt.Println("Bilangan Prima 		:", *data.Prima())
	fmt.Println("Hasil Genap / Ganjil 	:", *data.OddEven())
	fmt.Println("Bilangan Fibonacci 	:", *data.Fibonacci())
}

func (pri *Deret) Prima() *[]int {
	var temp []int
	for bilangan := 1; bilangan < pri.limit; bilangan++ {
		i := 0
		for bagi := 1; bagi < pri.limit; bagi++ {
			if bilangan%bagi == 0 {
				i++
			}
		}
		if (i == 2) && (bilangan != 1) {
			temp = append(temp, bilangan)
		}
	}
	return &temp
}

func (num *Deret) OddEven() *string {
	var result string
	if num.limit%2 == 0 {
		result = "Bilangan Genap"
	} else {
		result = "Bilangan Ganjil"
	}

	return &result
}

func (num *Deret) Fibonacci() *[]int {
	var result []int
	firstNumber := 0
	secondNumber := 1
	for i := 0; i < num.limit; i++ {
		fib := firstNumber
		firstNumber = secondNumber
		secondNumber = firstNumber + fib
		if fib <= num.limit {
			result = append(result, fib)
		}
	}
	return &result
}
