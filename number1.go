package main

import (
	"fmt"
)


func isPolycarpValidationNumber(num int) bool {


	if num % 3 == 0 || num % 10 == 3 {
		return false
	}
	return true
}


func findKthElemen(k int) int{
	jumlah := 0
	angka := 0

	for{
		angka++

		if isPolycarpValidationNumber(angka){
			jumlah++
		}

		if jumlah == k {
			break
		}
	}

	return angka
}


func main(){
	var t int
	fmt.Println("Input: ")
	fmt.Scan(&t)

	fmt.Println("-----------------------------------")
	fmt.Println("Output")
	for i := 0; i < t; i++ {
		var k int
		fmt.Scan(&k)
		result := findKthElemen(k)
		fmt.Println(result)
	}
}

