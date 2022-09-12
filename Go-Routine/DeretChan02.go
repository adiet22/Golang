package main

import "fmt"

func Fibonacci(num int) {

	wg.Add(2)
	c := make(chan int, num)

	go fibon(c)
	go OddEven(c)

	wg.Wait()

}

func fibon(c chan<- int) {
	defer func() {
		close(c)
		wg.Done()
	}()

	x, y := 0, 1
	for i := 0; i < cap(c); i++ {
		c <- x
		x, y = y, x+y
	}
}

func OddEven(c <-chan int) {
	defer func() {
		wg.Done()
	}()

	for i := range c {
		mtx.RLock()
		if i <= cap(c) {
			if i%2 == 0 {
				fmt.Println("Bilangan Genap", i)
			} else {
				fmt.Println("Bilangan Ganjil", i)
			}
		}
		mtx.RUnlock()
	}
}
