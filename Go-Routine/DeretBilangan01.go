package main

import "fmt"

type Deret struct {
	Limit int
}

func DeretBilangan(Angka int) {
	var data = &Deret{Angka}
	wg.Add(3)

	fmt.Println("\nLimit n			:", Angka)
	go data.Prima()
	go data.Fibonacci()
	go data.EvenOdd()

	wg.Wait()
}

func (num *Deret) Prima() {
	var result []int

	for bilangan := 1; bilangan < num.Limit; bilangan++ {
		i := 0
		for pembagi := 1; pembagi < num.Limit; pembagi++ {
			if bilangan%pembagi == 0 {
				i++
			}
		}
		if i == 2 && bilangan != 1 {
			result = append(result, bilangan)
		}
	}

	fmt.Println("Bilangan Prima 		:", result)
	wg.Done()
}

func (num *Deret) EvenOdd() {
	var result string

	if num.Limit%2 == 0 {
		result = "Bilangan Ganjil"
	} else {
		result = "Bilangan Genap"
	}
	fmt.Println("Hasil Genap / Ganjil 	:", result)
	wg.Done()
}

func (num *Deret) Fibonacci() {
	var result []int
	firstNumber := 0
	secondNumber := 1
	for i := 0; i < num.Limit; i++ {
		fib := firstNumber
		firstNumber = secondNumber
		secondNumber = firstNumber + fib
		if fib <= num.Limit {
			result = append(result, fib)
		}
	}
	fmt.Println("Bilangan Fibonacci 	:", result)
	wg.Done()

}
