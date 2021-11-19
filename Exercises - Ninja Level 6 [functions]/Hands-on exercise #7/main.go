package main

import "fmt"

// Assign a func to a variable, then call that func

func main(){

	anonymous := func(){
		fmt.Println("This is also an anonymous funciont, but now in a var")
	}

	anonymous()

}