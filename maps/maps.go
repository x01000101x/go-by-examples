package main

import "fmt"

func main() {

	//create map
	m := make(map[string]int)

	//assign map
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	//to get the value of the map specific
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	//len
	fmt.Println(len(m))

	//delete specific map by key
	delete(m, "k2")
	fmt.Println("delete:", m)

	//check if a specific map exist by key
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	//create map
	n := map[string]int{
		"foo": 1,
		"bar": 2,
	}
	fmt.Println("map:", n)
}
