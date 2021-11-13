package main

import "fmt"

// Using a COMPOSITE LITERAL:
// create a SLICE of TYPE int
// assign 10 VALUES
// Range over the slice and print the values out.
// Using format printing
// print out the TYPE of the slice

func main() {
	s := []int{9, 98, 87, 76, 65, 54, 43, 32, 21, 12}
	for i, j := range s {
		fmt.Println(i, j)
	}

	fmt.Printf("%T", s)
}
