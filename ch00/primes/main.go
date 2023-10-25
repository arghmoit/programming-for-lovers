package main

import (
	"fmt"
	"math"
	"time"
	"log"
)

func main() {
	fmt.Println("Finding Primes!")

	n := 100000000
	start := time.Now()
	TrivialPrimeFinder(n)
	elapsed := time.Since(start)
	log.Printf("TrivialPrimeFinder took %s", elapsed)

	start2 := time.Now()
	SieveOfEratosthenes(n)
	elapsed2 := time.Since(start2)
	log.Printf("SieveOfEratosthenes took %s", elapsed2)

	// fmt.Println(ListPrimes(11))
}

// TrivialPrimeFinder takes as input an integer n and return a slice of
// boolean variables storing the primality of each non-negative integer
// upto and including n
func TrivialPrimeFinder(n int) []bool {
	primeBooleans := make([]bool, n+1) // default false values
	
	for p := 2; p <= n; p++ {
		primeBooleans[p] = IsPrime(p)
	}
	
	return primeBooleans
}

// IsPrime takes as input an integer p and returns true if p is prime
// and false otherwise
func IsPrime(p int) bool {
	if p == 0 || p == 1 {
		return false
	}
	
	for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
		// is k a divisor of p? Is so, return false
		if p % k == 0 {
			return false		// composite
		}
	}
	
	return true
}

// SieveOfEratosthenes takes as input an integer n and returns a slice of
// boolean variables storing the primality of each non-negative integer
// upto and including n, implementing the "Sieve of Eratosthenes" algorithm.
func SieveOfEratosthenes(n int) []bool {
	primeBooleans := make([]bool, n+1)
	
	// everyone starts as false by default
	// now, set everything from 2 onward equal to true
	for p := 2; p <= n; p++ {
		primeBooleans[p] = true
	}
	
	for p := 2; float64(p) <= math.Sqrt(float64(n)); p++ {
		// if p prime? Is so, cross of its multiples
		if primeBooleans[p] == true {
			primeBooleans = CrossOfMultiples(primeBooleans, p)
		}
	}
	
	return primeBooleans
}

// CrossofMultiples takes as input a boolean slice primeBooleans and an
// integer p and returns an updated slice in which all variables in the array
// whose indices are multiples of p (greater than p) have been set to false.
func CrossOfMultiples(primeBooleans []bool, p int) []bool {
	n := len(primeBooleans) - 1

	// consider every multiple of p, starting with 2p, and "cross it off"
	// by setting its corresponding entry of the slice to false
	for k := 2*p; k <= n; k += p {
		primeBooleans[k] = false
	}
	return primeBooleans
}

// ListPrimes takes as input an integer n and returns a slice
// containing all prime numbers up to and (possibly) including n.
func ListPrimes(n int) []int {
	// first, create a slice of length 0
	primes := make([]int, 0)
	
	primeBooleans := SieveOfEratosthenes(n)
	for p := range primeBooleans {
		if primeBooleans[p] == true {
			// append p to our list
			primes = append(primes, p)
		}
	}
	
	return primes
}

/*
    TrivialPrimeFinder(n)
        primeBooleans <- array of length n + 1 false boolean variables
	    for every integer p from 2 to n
	        if IsPrime(p) is true
	            primeBooleans[p] <- true
	    return primeBooleans

     SieveOfEratosthenes(n)
        primeBooleans <- array of n + 1 true boolean variables
	    primeBooleans[0] <- false
	    primeBooleans[1] <- false
	    for every integer p between 2 and square root of n
	        if primeBooleans[p] = true
	            primeBooleans <- CrossOfMultiples(primeBooleans, p)
	    return primeBooleans

     CrossOfMultiples(primeBooleans, p)
         n <- length(primeBooleans) - 1
         for every multiple k of p (from 2p to n)
             primeBooleans[k] <- false
         return primeBooleans

     ListPrimes(n)
         primes <- array of integers of length 0
	     primeBooleans <- SieveOfEratosthenes(n)
	     for every integer p from 0 to n+1
	         if primeBooleans[p] = true
	             primes <- append(primes, p)
	     return primes
*/

