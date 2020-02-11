package main

import (
	"fmt"
	"math"
)

func generateSeive(maxN int) []int {

	// Create a seive of boolean values upto maxN+1
	seive := make([]bool, maxN+1)

	// Since by default all empty bool values are false
	// we will intialise all seive elements to be true
	// ie, all elements in the sieve are prime
	for i := 1; i < maxN+1; i++ {
		seive[i] = true
	}

	// If a number is prime (seive[i] = true) its corresponding
	// multiples are marked as not prime (seive[i] = false)
	for i := 2; i <= int(math.Sqrt(float64(maxN+1))); i++ {
		if seive[i] == true {
			for j := i * 2; j < maxN+1; j += i {
				seive[j] = false
			}
		}
	}

	// The remaining elements which are true are indeces
	// that are prime. We will now add all indeces to the
	// prime number list
	primeList := make([]int, 0)
	for i := 0; i < maxN+1; i++ {
		if seive[i] == true {
			primeList = append(primeList, i)
		}
	}

	return primeList
}

func main() {
	var inputN int
	var primes []int

	fmt.Printf("Enter the maximum number: ")
	fmt.Scanln(&inputN)

	primes = generateSeive(inputN)
	fmt.Println("Prime Numbers to ", inputN, ":", primes)
}
