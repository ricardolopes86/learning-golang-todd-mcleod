package main

import "fmt"

// Build and use an anonymous func

func main(){

	func (){
		fmt.Println("This is an anonymous function.")
	}()

}