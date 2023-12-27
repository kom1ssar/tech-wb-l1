package main

import (
	"fmt"
	"reflect"
)

// определям тип с помощью пакета reflect
func defineTypeByReflect(t interface{}) {
	// вспользуемся Kind, чтобы определять любой канал,как канал
	switch reflect.TypeOf(t).Kind() {

	case reflect.String:
		fmt.Println("string")
	case reflect.Int:
		fmt.Println("integer")
	case reflect.Bool:
		fmt.Println("boolean")
	case reflect.Chan:
		fmt.Println("channel")
	default:
		fmt.Println("type out of range")

	}

}

// определаем тип
func defineType(t interface{}) {
	switch t.(type) {

	case string:
		fmt.Println("string")
	case int:
		fmt.Println("integer")
	case bool:
		fmt.Println("boolean")
	case chan int:
		fmt.Println("channel integer")
	case chan string:
		fmt.Println("channel string")
	default:
		fmt.Println("type out of range")

	}

}

func main() {
	var a interface{}

	a = make(chan string)
	defineType(a)

	a = "yess"
	defineType(a)

	a = true
	defineType(a)

	a = 1
	defineType(a)

	fmt.Println("\ndefine by Reflect:")
	var b interface{}

	b = make(chan string)
	defineTypeByReflect(b)

	b = "yess"
	defineTypeByReflect(b)

	b = true
	defineType(b)

	b = 1
	defineTypeByReflect(b)

}
