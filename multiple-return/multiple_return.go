package main

import "fmt"

func multipleReturn() (int, int) {
	return 3, 7
}

func main() {
	//getting both variables
	a, b := multipleReturn()
	fmt.Println(a, b)

	//only getting the second variable
	_, c := multipleReturn()
	fmt.Println(c)
}
