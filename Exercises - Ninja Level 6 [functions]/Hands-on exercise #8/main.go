package main

import "fmt"

// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

func foo() func() int {
	return func() int {
		return 50
	}
}

func main() {

	f := foo()
	fmt.Println(f())

}
