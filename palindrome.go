package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	str = strings.ToLower(str)
	length := len(str)

	for i := 0; i < length/2; i++ {
		if str[i] != str[length-i-1] {
			return false
		}
	}

	return true
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)

		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {

	var kata string
	fmt.Println("Input: ")
	fmt.Scan(&kata)

	if(isPalindrome(kata) == true){
		fmt.Println(kata, "itu palindrome:")	
	}else{
		fmt.Println(kata, "bukan palindrome:")
	}
}
