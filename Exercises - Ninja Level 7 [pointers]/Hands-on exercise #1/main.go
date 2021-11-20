package main

import "fmt"

// Create a value and assign it to a variable.
// Print the address of that value.


func main(){
	var p int
	p = 43

	fmt.Printf("address of variable p is %v", &p)
}