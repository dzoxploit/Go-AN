package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate slice of random numbers
	numbers := generateRandomNumbers(10)

	fmt.Println("Sebelum pengurutan:")
	fmt.Println(numbers)

	// Sort the numbers using quicksort algorithm
	quicksort(numbers, 0, len(numbers)-1)

	fmt.Println("Setelah pengurutan:")
	fmt.Println(numbers)
}

// Menghasilkan slice bilangan acak
func generateRandomNumbers(n int) []int {
	numbers := make([]int, n)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

// Algoritma quicksort untuk mengurutkan slice bilangan
func quicksort(numbers []int, low, high int) {
	if low < high {
		// Membagi slice menjadi dua bagian dan mengambil pivot
		pivotIndex := partition(numbers, low, high)

		// Rekursif melakukan pengurutan pada kedua bagian
		quicksort(numbers, low, pivotIndex-1)
		quicksort(numbers, pivotIndex+1, high)
	}
}

// Mengembalikan pivot dan mengatur elemen-elemen yang lebih kecil di sebelah kiri
// elemen-elemen yang lebih besar di sebelah kanan
func partition(numbers []int, low, high int) int {
	pivot := numbers[high]
	i := low - 1

	for j := low; j < high; j++ {
		if numbers[j] < pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}
