package main

import "fmt"

func main() {

	a := 1
	b := 2

	// самый простой и очевидный вариант
	a, b = b, a
	fmt.Println(a)
	fmt.Println(b)

	// побитовый обмен (обмен с помощью xor)
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println(a)
	fmt.Println(b)
}
