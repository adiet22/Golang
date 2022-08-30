package main

import (
	"math/rand"
	"strconv"
	"strings"
)

func genPass(password, level string) string {
	var result string
	lev := strings.ToLower(level)

	if len(password) > 0 {

		//Apabila password kurang dari 6 krakter
		if len(password) <= 6 {
			if lev == "weak" {
				temp := password + ranWord()
				result = toUpper(temp)
			} else if lev == "middle" {
				temp := password + ranWord()
				result = toUpper(temp) + ranSymbol()
			} else if lev == "strong" {
				temp := password + ranWord()
				result = toUpper(temp) + ranSymbol() + ranNumb()
			} else {
				result = "Wrong Input Level !!"
			}

			//Apabila password lebih dari 6 karakter
		} else {
			if lev == "weak" {
				result = toUpper(password)
			} else if lev == "middle" {
				result = toUpper(password) + ranSymbol()
			} else if lev == "strong" {
				result = toUpper(password) + ranSymbol() + ranNumb()
			} else {
				result = "Wrong Input Level !!"
			}
		}
	} else {
		result = "Password Cant be Empty !!"
	}
	return result
}

func random(angka int) string {
	acak := rand.Intn(angka)
	result := strconv.Itoa(acak)
	return result
}

func ranSymbol() string {
	specialChar := "!@#$%^&*(){}<>"
	acak := rand.Intn(14)
	temp := string(specialChar[acak])
	return temp
}

func toUpper(password string) string {
	depan := strings.ToUpper(string(password[0]))
	tengah := password[1 : len(password)-1]
	belakang := strings.ToUpper(string(password[len(password)-1]))
	result := depan + tengah + belakang
	return result
}

func ranAbjad() string {
	abjad := "abcdefghijklmnopqrstuvwxyz"
	acak := rand.Intn(25)
	var temp = string(abjad[acak])
	return temp
}

func ranWord() string {
	var a []string

	for len(a) <= 5 {
		a = append(a, ranAbjad())
	}
	b := strings.Join(a, "")

	return b
}

func ranNumb() string {
	var c []string
	for len(c) <= 2 {
		c = append(c, random(9))
	}
	d := strings.Join(c, "")
	return d
}
