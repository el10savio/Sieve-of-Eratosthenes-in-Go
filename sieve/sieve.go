package sieve

import (
	"math"
)

// GenerateSieve returns a list of prime numbers up to maxN
func GenerateSieve(maxN int) []int {
	// Create a sieve of boolean values upto maxN+1
	maxN++
	sieve := make([]bool, maxN)

	// Since by default all empty bool values are false
	// we will initialize all sieve elements to be true
	// ie, all elements in the sieve are prime
	for i := 1; i < maxN; i++ {
		sieve[i] = true
	}

	// If a number is prime (sieve[i] = true) its corresponding
	// multiples are marked as not prime (sieve[i] = false)
	for i := 2; i <= int(math.Sqrt(float64(maxN))); i++ {
		if sieve[i] {
			for j := i * 2; j < maxN; j += i {
				sieve[j] = false
			}
		}
	}

	// The remaining elements which are true are indeces
	// that are prime. We will now add all indeces to the
	// prime number list
	primeList := make([]int, 0)
	for i := 0; i < maxN; i++ {
		if sieve[i] {
			primeList = append(primeList, i)
		}
	}

	return primeList
}
