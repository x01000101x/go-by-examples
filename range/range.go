package main

import "fmt"

func main() {

	//Creating slice and summing all the values using for range
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sums", sum)

	//Getting index
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//Create map and for-looping the key and value
	kvs := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	for k, v := range kvs {
		fmt.Printf("%s : %s \n", k, v)
	}

	//only getting keys from the map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	//rune
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
