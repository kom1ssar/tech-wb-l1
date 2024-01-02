package main

import (
	"fmt"
	"strings"
)

func revers(str string) string {

	// созадем слайс из строки
	words := strings.Split(str, " ")

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {

		// меняем местами
		words[i], words[j] = words[j], words[i]

	}
	// джоиним слайс в строку
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(revers("snow dog sum"))
}
