package main

import "fmt"

// Create and use an anonymous struct.

func main(){

	person1 := struct {
		name string
		last_name string
		age int
	}{
		name: "Ayrton",
		last_name: "Senna",
		age: 34,
	}

	fmt.Println(person1)

}