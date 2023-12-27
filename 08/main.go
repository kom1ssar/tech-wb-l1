package main

import (
	"errors"
	"fmt"
)

// setBit меняем бит числа по укзателю.
//
//	num - укзаатель изменяемого числа.
//
// position - позиция изменяемого бита.
// bitValue -  устанавливаемое значение бита.
func setBit(num *int64, position uint, bitValue int8) error {

	// валидируем позицию бита
	if position > 63 {
		return errors.New("position out of range [0-63]")
	}

	// валидируем устанавливаемое значение бита
	if bitValue > 1 || bitValue < 0 {
		return errors.New("bitValue out of range [0-1]")

	}

	if bitValue == 1 {
		// устанавливаем бит по индексу position на 1
		*num = *num | (1 << position)
	} else {
		// устанавливаем бит по индексу position на 0
		*num = *num &^ (1 << position)
	}

	return nil

}

func main() {

	// 010
	var firstNum int64 = 2

	// firstNum = 2  -  010  меняем бит по индексу 2 на единицу = 110 - получем 6
	err := setBit(&firstNum, 2, 1)
	if err != nil {
		fmt.Println(err)
	}

	// 110
	fmt.Println(firstNum)

	// 111
	var secondNum int64 = 15

	// firstNum = 15  -  1111  меняем бит по индексу 2 на ноль = 1011 - получем 11
	err = setBit(&secondNum, 2, 0)
	if err != nil {
		fmt.Println(err)
	}

	// 1011
	fmt.Println(secondNum)

}
