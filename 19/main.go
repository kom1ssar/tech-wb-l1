package main

import (
	"fmt"
)

func stringRevers(str string) string {
	// создаем слайс рун из строки
	runes := []rune(str)

	// будем бежать по слайсу двумя поинтерами ( с начала и конца)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		// меняем местами руны
		runes[i], runes[j] = runes[j], runes[i]
	}

	// переводим слайс рун в строку
	return string(runes)

}

func main() {
	fmt.Println(stringRevers("главрыба"))
}
