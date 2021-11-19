package main

import "fmt"

// create a func with the identifier foo that
// takes in a variadic parameter of type int
// pass in a value of type []int into your func (unfurl the []int)
// returns the sum of all values of type int passed in
// create a func with the identifier bar that
// takes in a parameter of type []int
// returns the sum of all values of type int passed in

func foo(n ...int) int {
	var sum int
	for _, v := range n {
		sum = sum + v
	}
	return sum
}

func bar(n []int) int {
	var total int
	for _, v := range n {
		total += v
	}
	return total
}

func main() {

	n := []int{12, 34, 56, 78, 90}
	fmt.Printf("%d\n", foo(n...))

	x := []int{9,87, 65, 54, 32, 1}
	fmt.Printf("%d\n", foo(x...))

}
