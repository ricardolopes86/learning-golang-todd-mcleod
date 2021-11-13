package main

import "fmt"

// Create your own type “person” which will have an underlying type of “struct”
// so that it can store the following data:
// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values,
// ranging over the elements in the slice which stores the favorite flavors.

type person struct{
	first_name string
	last_name string
	favorite_icecream []string
}

func main(){
	person1 := person{
		first_name: "Snoopy",
		last_name: "Dog",
		favorite_icecream: []string{"vanilla","leaver"},
	}

	person2 := person{
		first_name: "Tom",
		last_name: "Cat",
		favorite_icecream: []string{"Java","EJB"},
	}

	fmt.Println("Person:", person1.first_name, person1.last_name)
	fmt.Println("Favorite ice creams:")
	for _,k := range person1.favorite_icecream{
		fmt.Printf("%s", k)
		fmt.Println()
	}

	fmt.Println("Person:", person2.first_name, person2.last_name)
	fmt.Println("Favorite ice creams:")
	for _, v := range person2.favorite_icecream {
		fmt.Printf("%s", v)
		fmt.Println()
	}

}