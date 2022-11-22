package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("divisible by 4")
	}

	if i := 9; i < 0 {
		fmt.Println("i is minus")
	} else if i < 10 {
		fmt.Println("i only have 1 digit")
	} else {
		fmt.Println("i have multiple digits")
	}

}
