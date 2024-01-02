package main

import (
	"fmt"
	"strings"
)

func uniqueLetters(str string) bool {
	// делаем мапку
	hashTable := make(map[string]struct{})

	// бежим по строке
	for _, v := range strings.ToLower(str) {
		_, ok := hashTable[string(v)]
		// если есть в мапе возвращаем фолс
		if ok {
			return false
		}
		// если нет добавляем
		hashTable[string(v)] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(uniqueLetters("abC"))
}
