package main

import (
	"fmt"
)

// AddOne takes an integer m as input and returns the value of k+1.
func AddOne(m int) int {
	m = m + 1
	return m
}

// SumTwoInts takes two as input two integers and returns their sum
func SumTwoInts(a, b int) int {
	return a + b
}

// DoubleAndDuplicate takes as input a float64
// It returns two copies of that input variable.
func DoubleAndDuplicate(x float64) (float64, float64) {
	return 2.0 * x, 2.0 * x
}

// Pi takes no inputs and returns the value of pi,
// the mathematical constant
func Pi() float64 {
	return 3.14
}

// PrintHi simply prints a message "Hi" to the console.
func PrintHi() {
	fmt.Println("Hi")
}

func main() {
	fmt.Println("Functions!")

	x := 3 // x has type float
	n := SumTwoInts(x, x)
	fmt.Println(n)

	m := 17
	fmt.Println(AddOne(m))
	fmt.Println(m) // what gets printed?
}

