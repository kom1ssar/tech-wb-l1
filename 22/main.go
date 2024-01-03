package main

import (
	"fmt"
	"math/big"
)

func mul(a, b big.Int) *big.Int {

	return new(big.Int).Mul(&a, &b)
}

func div(a, b big.Int) *big.Int {

	return new(big.Int).Div(&a, &b)
}

func sub(a, b big.Int) *big.Int {

	return new(big.Int).Sub(&a, &b)
}

func sum(a, b big.Int) *big.Int {

	return new(big.Int).Add(&a, &b)
}

// a > 2 ^ 20 поэтому будет использовать big.Int и сделаем имплемитацию его встроенных методов
func main() {
	a, _ := new(big.Int).SetString("92448576", 10)
	b, _ := new(big.Int).SetString("1148576", 10)

	fmt.Printf("mul: %d\n", mul(*a, *b))
	fmt.Printf("div: %d\n", div(*a, *b))
	fmt.Printf("sub: %d\n", sub(*a, *b))
	fmt.Printf("sum: %d\n", sum(*a, *b))
}
