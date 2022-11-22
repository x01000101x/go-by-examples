package main

import "fmt"

func main() {
	//create empty slice
	s := make([]string, 3)
	fmt.Println(s)

	//setting and getting slice values
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set", s)
	fmt.Println("get", s[2])

	//len
	fmt.Println("len:", len(s))

	//adding values to the slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	//creating new slice and copying the values from the slice above
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	//slicing values from key 2 to 5
	l := s[2:5]
	fmt.Println("sl1: ", l)

	//slicing values from right
	l = s[:5]
	fmt.Println("sl2: ", l)

	//slicing values from left
	l = s[2:]
	fmt.Println("sl3:", l)

	//creating slice
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	//creating 2D slice
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := 1 + i
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
