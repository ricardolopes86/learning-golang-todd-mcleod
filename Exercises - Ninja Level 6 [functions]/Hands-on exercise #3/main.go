package main

import "fmt"

// Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

func foo(){
	fmt.Println("Defered function return")
}

func main(){
	defer foo()
	fmt.Println("This message is placed after foo() function")

}