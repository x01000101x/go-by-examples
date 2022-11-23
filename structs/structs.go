package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(person{"Bob", 22})
	fmt.Println(person{name: "Jackal", age: 18})
	fmt.Println(person{name: "Jackal", age: 18})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Sewey", age: 19})
	fmt.Println(newPerson("Choy"))

	s := person{name: "Sean", age: 19}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

}
