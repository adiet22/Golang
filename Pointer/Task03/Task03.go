package Task03

import (
	"math"
)

type Permukaan struct {
	Sisi float64
}

func (num *Permukaan) Luas() float64 {
	var result float64
	result = math.Pow(num.Sisi, 2)
	return result
}

func (num *Permukaan) Keliling() float64 {
	var result float64
	result = 4 * num.Sisi
	return result
}

func (num *Permukaan) Volume() float64 {
	var result float64
	result = math.Pow(num.Sisi, 3)
	return result
}
