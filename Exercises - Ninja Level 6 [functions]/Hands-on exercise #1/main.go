package main

import "fmt"

// Hands on exercise
// create a func with the identifier foo that returns an int
// create a func with the identifier bar that returns an int and a string
// call both funcs
// print out their results

func foo() int {
	return 35
}

func bar() string {
	return "returns a string"
}

func main(){
	fmt.Printf("%d\n",foo())
	fmt.Printf(bar())
}