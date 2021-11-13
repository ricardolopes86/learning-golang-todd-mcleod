package main

import "fmt"

// Using the code from the previous example, use SLICING to create the following new slices which are then printed:
// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]

func main() {

	s := []int{9, 98, 87, 76, 65, 54, 43, 32, 21, 12}
	s1 := s[:5]
	s2 := s[1:6]
	s3 := s[5:]
	s4 := s[3:8]

	fmt.Println(s1)

	fmt.Println(s2)

	fmt.Println(s3)

	fmt.Println(s4)

}
