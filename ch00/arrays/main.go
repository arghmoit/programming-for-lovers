package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays (and slices).")

	var list [6]int
	
	list[0] = -8

	i := 3
	list[2*i-4] = 17

	list[len(list)-1] = 43

	// these commands lead to compiler errors
	// list[len(list)] = 7
	// list[-4] = 2
	
	fmt.Println(list)

	// slice declarations are a little different
	var a []int				  // right now,  a has a special value nil
	n := 4

	// slices must be made
	a = make([]int, 4)

	// we set values of arrays similarly to those of slices
	a[0] = 6
	a[2] = -33
	
	fmt.Println(a)

	// one-line declaration are used in practice
	b := make([]int, n+2)
	fmt.Println(b)

	fmt.Println(FactorialArray(6))

	var c [6]int
	d := make([]int, 6)
	ChangeFirstElementArray(c)
	ChangeFirstElementSlice(d)
	fmt.Println("c is now", c)
	fmt.Println("d is now", d)

	e := make([]int, 5)
	e[0] = 1
	e[1] = 0
	e[2] = -23
	e[3] = 45
	e[4] = -240
	fmt.Println("minimum value is", MinIntegerArray(e))

	fmt.Println(MinIntegers(3, 6, 7, 1, 5, -43))
}

// variadic functions take a variable number of inputs

// MinIntegers takes as input an arbitrary number of integers
// and returns their minimum value
func MinIntegers(numbers ...int) int {
	if len(numbers) == 0 {
		panic("Error: empty slice given to MinIntegers().")
	}

	return MinIntegerArray(numbers)
}

// MinIntegerArray takes as input a slice of integers and returns
// the minimum value in that slice
func MinIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("Error: empty slice given to MinIntegerArray().")
	}

	var minVal int			// will store the minimum value

	// range over list, and if we are at the 0th element or we find
	// something smaller than m, update m
	// for p := 0; p < len(list); p++ {
	// 	if p == 0 || minVal > list[p] {
	// 		minVal = list[p]
	// 	}
	// }

	// for p := range list {		// this is equivalent to
	// 	// for i := 0; i < len(list; i++
	// 	if p == 0 || minVal > list[p] {
	// 		minVal = list[p]
	// 	}
	// }

	for p, val := range list {
		if p == 0 || minVal > val {
			minVal = val
		}
	}

	return minVal
}

func ChangeFirstElementArray(a [6]int) {
	a[0] = 1					// copy of input array gets edited
	// copy is destroyed
}

func ChangeFirstElementSlice(a []int) {
	a[0] = 1
	// when you pass in a slice to a function, you get access to
	// that underlying array
}

// FactorialArray takes as input an integer n, and it returns
// an array of length n+1 whose kth element is k!
func FactorialArray(n int) []int {
	if n < 0 {
		panic("Error: non-positive input given to FactorialArray().")
	}

	fact := make([]int, n+1)

	// recall: 0! = 1
	fact[0] = 1

	// range k from 1 to n, and use the fact that k! = k * (k-1)!
	for k := 1; k <= n; k++ {
		fact[k] = fact[k-1] * k
	}

	return fact
}

    // FactorialArray(n)
    //     a <- array of length n
	// a[0] <- 1
	// for every integer k between 1 and n
	//     a[k] <- a[k-1] * k
	// return a

