package main

import "fmt"

// Using a COMPOSITE LITERAL:
// create an ARRAY which holds 5 VALUES of TYPE int
// assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing
// print out the TYPE of the array

func main() {
	a := [5]int{12, 23, 34, 45, 56}

	for i, j := range a {
		fmt.Println(i, j)
	}

	fmt.Printf("%T", a)
}
