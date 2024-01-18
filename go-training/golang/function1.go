package main

import (
	"fmt"
	"math"
)

type Operator func(x float64) float64

// Map applies op to each element of a.
func Map(op Operator, a []float64) []float64 {
	res := make([]float64, len(a))
	for i, x := range a {
		res[i] = op(x)
	}
	return res
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Total Nums: ", total)
}

func calls() (int, int) {
	return 10, 11
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	sum(1, 2, 3)
	sum(4, 5, 6)

	func(a int, b int) {
		c := a + b
		fmt.Println("This is anonymous function, total", c)
	}(5, 5)

	nums := []int{3, 3, 3}
	sum(nums...)

	a, b := calls()
	fmt.Println("Result a, ", a)
	fmt.Println("Result b, ", b)

	_, c := calls()
	fmt.Println("Result c, ", c)

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	op := math.Abs
	a := []float64{1, -2}
	b := Map(op, a)
	fmt.Println(b) // [1 2]

	c := Map(func(x float64) float64 { return 10 * x }, b)
	fmt.Println(c) // [10, 20]
}
