package main

import "strconv"

func menonton(data []int, n int) string {
	var result string
	var temp string
	for index, _ := range data {
		if index == 0 {
			result = "Data is only one !!"
		} else if index > 0 {
			temp = sum(data, n)
			result = temp
			break
		} else {
			result = "Data is empty !!"
		}
	}
	return result
}

func sum(data []int, n int) string {
	var result string
	var temp string

	for i := 0; i < len(data)-1; i++ {
		for j := len(data) - 1; j >= i; j-- {
			if data[i]+data[j] == n && data[i] != data[j] {
				temp = strconv.Itoa(data[i]) + " and " + strconv.Itoa(data[j])
				return temp
			} else {
				result = "Data not found !!"
			}
		}
	}
	return result
}
