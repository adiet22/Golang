package main

import (
	"fmt"
	"regexp"
)

func CheckPass(data string) {
	wg.Add(2)
	c := make(chan string, 2)

	go Checking(c)
	go Process(data, c)

	wg.Wait()
}

func Process(pass string, c chan<- string) {
	defer func() {
		close(c)
		wg.Done()
	}()

	var res1, res2, res3, res4, res5 string

	r, _ := regexp.Compile("[a-z]+")
	if r.MatchString(pass) != true {
		res1 = "Password Harus mengandung Alfabet"
	}
	c <- res1

	r2, _ := regexp.Compile("[A-Z]+")
	if len(r2.FindStringSubmatchIndex(pass)) < 2 {
		res2 = "Password Harus mengandung Kapital Minimal 1"
	}
	c <- res2

	r3, _ := regexp.Compile("[!@#$%^&*()_+~`{}.,/?><|]+")
	if r3.MatchString(pass) != true {
		res3 = "Password Harus Berisi Simbol Minimal 1"
	}
	c <- res3

	r4, _ := regexp.Compile("[0-9]+")
	if r4.MatchString(pass) != true {
		res4 = "Password Harus Berisi Angka Minimal 1"
	}
	c <- res4

	if len(pass) < 8 {
		res5 = "Panjang Password Minimal 8 Karakter"
	}
	c <- res5

}

func Checking(c <-chan string) {
	defer func() {
		mtx.RUnlock()
		wg.Done()
	}()

	var value []string
	mtx.RLock()
	for v := range c {
		if len(v) > 0 {
			value = append(value, v)
		}
	}
	if len(value) <= 0 {
		fmt.Println("\n Password Berhasil Dibuat")
	} else {
		fmt.Println("\n Error Checking Password :")
		for _, i := range value {
			fmt.Printf("%+q \n", i)
		}
	}
}
