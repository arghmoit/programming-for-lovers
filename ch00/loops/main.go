package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops!")

	fmt.Println(Factorial(5))
	fmt.Println(Factorial(300))

	i := 0
	n := 25
	for i <= n {
		fmt.Println("At", i, "Factorial(", i, ") is", Factorial(i))
		i = i + 1
	}

	fmt.Println(SumFirstNIntegers(10))

	fmt.Println(AnotherSum(10))

	fmt.Println(GaussSum(10))

	fmt.Println(SumEven(10))

	var count uint = 10
	for ; count >= 0; count-- {
		fmt.Println(count)
	}
	fmt.Println("Blast off!")
}

// Factorial takes as input an integer n and returns n!
// n! = n * (n-1) * (n-2) * ... * 2 * 1
//
// Factorial(n)
//     p <- 1
//     i <- 1
//     while i <= n
//         p <- p * i
//         i <- i + 1
//     return p
//
// mathematical fact: n! = n * (n-1)!
// when n = 1: 1! = 1 * 0!
// so, 1 = 1 * 0! and therefore 1 = 0!

func Factorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial().")
	}

	p := 1 // will store product
	i := 1 // serves as a counter

	// Go has no keyword "while" and uses "for" instead
	for i <= n {
		p = p * i
		i = i + 1
	}

	fmt.Println("i is currently", i)
	
	return p
}

// AnotherFactorial(n)
//     p <- 1
//     for every integer i between 1 and n
//         p <- p * i
//     return p
func AnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial().")
	}
	
	p := 1

	for i := 1; i <= n; i++ {
		// i := 1 is called the "initialization step"
		// i <= n is called the "condition"
		// i++ is called the "postcondition"
		p = p * i
	}

	// fmt.Println("i is currently", i)
	
	return p
}

func YetAnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial().")
	}

	p := 1

	for i := n; i >= 1; i-- {
		p *= i
	}

	return p
}

// SumFirstNIntegers takes as input an integer n
// and returns the sum of the first n positive integers
func SumFirstNIntegers(n int) int {
	if n < 0 {
		panic("Error: non-positive argument given to SumFirstNIntegers().")
	}
	
	sum := 0					// will store the final answer
	// also have sum -= i, sum *= i, and sum /= i

	i := 1						// represents current integer being added to sum
	for i <= n {
		sum += i				// equivalent to sum = sum + i
		i++						// there is also i-- for i = i - 1
	}

	return sum
}

// AnotherSum takes an integer n as input and returns the sum
// of the first n positive integers, using a for loop
func AnotherSum(n int) int {
	if n < 0 {
		panic("Error: non-positive argument given to AnotherSum().")
	}

	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum
}

func GaussSum(n int) int {
	return n * (n+1) / 2
}

// SumEven takes as input an integer k and returns the sum
// of all even positive integers up to and (possibly) including k.
func SumEven(n int) int {
	if n < 0 {
		panic("Error: non-positive argument given to SumEven().")
	}

	sum := 0

	// (one way) for every even integer i between 2 and k, add i to sum
	for i := 2; i <= n; i += 2 {
		// i must be even :)
		sum += i
	}

	// (another way) for every even integer i between 2 and k, add i to sum
	// for i := 2; i < n; i++ {
	// 	// is i even?
	// 	if i % 2 == 0 {
	// 		// yes!
	// 		sum += i
	// 	}
	// }

	return sum
}

