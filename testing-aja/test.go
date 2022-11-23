package main

import "fmt"

type orang struct {
	tinggi int
	nama   string
}

func main() {
	// 	n := 12
	// 	var r *int = &n

	// 	n = 7

	// 	fmt.Println("N", n)
	// 	fmt.Println("&N", &n)
	// 	fmt.Println("R", r)
	// 	fmt.Println("*R", *r)

	// }

	// 	var number = 4
	// 	fmt.Println("before :", number) // 4

	// 	change(&number, 10)
	// 	fmt.Println("after  :", number) // 10
	// }

	// func change(original *int, value int) {
	// 	*original = value

	x := orang{tinggi: 173, nama: "Leo"}
	fmt.Println(x)

	var w *orang = &x
	fmt.Println(w.nama)

	w.nama = "Rizki"
	fmt.Println(w.nama)
	fmt.Println(x.nama)

}
