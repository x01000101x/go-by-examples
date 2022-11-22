package main

import "fmt"

//the variadic functions
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)

}

func main() {

	//calling functions
	sum(1, 2)
	sum(88, 91, 24, 51, 0, 32)

	//slice (if u are using slice add ...)
	numbers := []int{1, 2, 3, 4}
	sum(numbers...)
}
