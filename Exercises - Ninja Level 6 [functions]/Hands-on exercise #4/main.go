package main

import "fmt"

// Create a user defined struct with
// the identifier “person”
// the fields:
// first
// last
// age
// attach a method to type person with
// the identifier “speak”
// the method should have the person say their name and age
// create a value of type person
// call the method from the value of type person

type Person struct {
	first string
	last  string
	age   int
}

func (s Person) speak() {
	fmt.Println(s.first)
}

func main() {

	me := Person{
		first: "Peter",
		last: "Parker",
		age: 18,
	}

	me.speak()

}
