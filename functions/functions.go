package main

import "fmt"

//creating functions

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a int, b int, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println("The result of plus :", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("The result of plusPlus :", res)
}
