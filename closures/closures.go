package main

import "fmt"

//the function returns closure (function that dont have name)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	//calling the closure
	newInt := intSeq()
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())

	newInts := intSeq()
	fmt.Println(newInts())

}
