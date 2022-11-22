package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial :", i)

	//pemanggilan func biasa
	zeroval(i)
	fmt.Println("zeroval :", i)

	//pemanggilan menggunakan pointer, hasilnya kebalikan dari yg biasa
	zeroptr(&i)
	fmt.Println("zeroptr :", i)

	fmt.Println("Pointer :", i)
}
