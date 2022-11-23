package main

import (
	"errors"
	"fmt"
)

func test(a int, b int) (int, error) {
	if a/b == 10 {
		return 0, errors.New("error tidak bisa 0")
	} else {
		c := a / b
		return c, nil
	}
}

func main() {
	hasil, err := test(100, 10)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
