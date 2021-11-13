package main

// Create a slice of a slice of string ([][]string). 
// Store the following data in the multi-dimensional slice:

// "James", "Bond", "Shaken, not stirred"
// "Miss", "Moneypenny", "Helloooooo, James."

// Range over the records, then range over the data in each record.


import "fmt"

func main(){

	names1 := []string{
		"James","Bond","Shaken, not stirred",
	}
	names2 := []string{
		"Miss","Moneypenny","Helloooooo, James.",
	}

	combined := [][]string{names1,names2}

	for i, x := range combined {
		fmt.Println("Record: ", i)
		for _, j := range x {
			fmt.Println(j)
		}
	}
}