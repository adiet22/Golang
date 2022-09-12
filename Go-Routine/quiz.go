package main

import "fmt"

func sum(d []int, ch chan int) {
	sum := 0
	for _, v := range d {
		//hitung
		sum += v
		//fmt.Print(v)
	}
	//send result to result
	ch <- sum // send sum to c
}

func hitung() {
	a := []int{7, 10, 2, 34, 33, -12, -8, 4}
	chn := make(chan int)
	go sum(a[:len(a)/2], chn)
	go sum(a[len(a)/2:], chn)
	//receive result from channel and print
	x, y := <-chn, <-chn
	// receive from c
	//fmt.Println("\n", x+y)
	fmt.Println("\n", x)
	fmt.Println("\n", y)

	//fmt.Println(y)
}
