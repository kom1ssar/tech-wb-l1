package main

import "fmt"

// удаляем с сохранением порядка (более медленно т.к. делаем сдвиг элементов)
func remove(arr []int, i int) []int {

	copy(arr[i:], arr[i+1:]) //  сдвигаем элементы
	return arr[:len(arr)-1]
}

// быстрй способ, но порядок не сохраняется (меняем удаляемый элемент с последним и срезаем)
func fastRemove(arr []int, i int) []int {

	arr[i] = arr[len(arr)-1] // сохраняем последний элемент по индексу удаляемого элемента
	return arr[:len(arr)-1]
}

func main() {
	arr := []int{1, 2, 34, 5, 6, 7, 7, 8, 89}

	arr = remove(arr, 0)
	fmt.Println(arr)

	arr = fastRemove(arr, 0)
	fmt.Println(arr)

}
