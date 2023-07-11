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
