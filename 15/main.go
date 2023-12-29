package main

import (
	"fmt"
	"strings"
)

//var justString string
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100]
//}
//
//func main() {
//	someFunc()
//}

// Судя по всему функция  "createHugeString" - возвращает строку по переданому размеру.
// 1 << 10  побитовое смещение  = 1024
// Основная проблема,как мне кажется - это в создание строки по срезу в этой строке "justString = v[:100]" (строка фактически является слайсом, поэтому будет ссылка на копию базового слайса)
// После выхода из функции "someFunc" созданная строка длинной 1024 никуда не денется из памяти, потому что на нее ссылается
// строка  который хранится в глобальной перменной "justString", что может привести к утечки памяти
// самый простой вариант решения проблемы -  воспользоваться методом "Clone" из пакета  "strings"

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])

}

func main() {
	someFunc()
	fmt.Println(justString)
}
