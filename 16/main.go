package main

import "fmt"

func partition(arr []int, low int, high int) ([]int, int) {
	// опороное число
	pivot := arr[high]
	// нужна для разделения массива на меньше и большее
	wall := low

	//  начинаем итерироваться с  "wall" - нашего разделителья
	for i := wall; i < high; i++ {
		// если число меньше опрного
		if arr[i] < pivot {
			// тогда меяем это чисило с разделителем
			arr[i], arr[wall] = arr[wall], arr[i]
			// смещаем разделитель
			wall++
		}
	}

	// менняем места разделительное число с опорным
	arr[wall], arr[high] = arr[high], arr[wall]

	// возвращаем массив и место индекс разделителя
	return arr, wall
}

func quick(arr []int, low int, high int) []int {
	// проверяем, что low не пересек high
	if low < high {
		// вызываем нашу функцию сортирвки "partition"
		arr, i := partition(arr, low, high)
		// рекурсивно вызываемся смещая индекс high на - 1
		arr = quick(arr, low, i-1)
		// рекурсивно вызываемся смещая индекс low на + 1
		arr = quick(arr, i+1, high)
	}

	// возвращаем массив
	return arr
}

// входная точка в  сортировку
func quickSort(arr []int) []int {
	return quick(arr, 0, len(arr)-1)
}

func main() {
	a := []int{122, 4, 23, 6, 43, 8, 9, 10, 22}
	fmt.Println(a)
	fmt.Println(quickSort(a))
}
