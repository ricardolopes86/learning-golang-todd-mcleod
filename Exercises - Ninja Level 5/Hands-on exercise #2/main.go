package main

import (
	"fmt"
)

// Take the code from the previous exercise,
// then store the values of type person in a map with the key of last name.
// Access each value in the map. Print out the values, ranging over the slice.

type person struct{
	first_name string
	last_name string
	favorite_icecream []string
}

func main(){

	p1 := person{
		first_name: "Bruce",
		last_name: "Benner",
		favorite_icecream: []string{
			"Peas",
			"Olive",
			"Spinach",
		},
	}

	p2 := person{
		first_name: "Peter",
		last_name: "Parker",
		favorite_icecream: []string{
			"Mosquito",
			"Flies",
		},
	}

	persons := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for j,k := range persons{
		fmt.Println(j,k.first_name)
		for _,v := range k.favorite_icecream{
			fmt.Println(v)
		}
	}

}