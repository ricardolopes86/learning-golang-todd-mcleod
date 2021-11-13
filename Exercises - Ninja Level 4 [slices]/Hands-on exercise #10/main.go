package main

// Using the code from the previous example, delete a record from your map. 
// Now print the map out using the “range” loop

import "fmt"

func main(){
	person := map[string][]string{
		"bond_james":{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss":{"James Bond", "Literature", "Computer Science"},
		"no_dr":{"Being evil", "Ice cream", "Sunsets"},
	}

	person["silva"] = append(person["silva"], "Classical Music", "Masks", "Sciense")

	for k, v := range person {
		fmt.Println(k)
		for i, j := range v{
			fmt.Printf("Position: %v, Value: %v\n", i, j)
		}
	}

	delete(person, "silva")
	fmt.Println("Now without latest record.")
	
	for k, v := range person {
		fmt.Println(k)
		for i, j := range v{
			fmt.Printf("Position: %v, Value: %v\n", i, j)
		}
	}

}