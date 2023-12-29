package main

import "fmt"

// бинарный поиск
func binarySearch(nums []int, target int) int {
	// начальный индекс массива
	left := 0
	// крайний иднекс массива
	right := len(nums) - 1

	// провряем не равняется ли первй элемент массива нашем таргетом
	if nums[left] == target {
		return left
	}

	// провряем не равняется ли последий элемент массива нашем таргетом
	if nums[right] == target {
		return right
	}

	for left <= right {
		// находим середину массива
		mid := (left + right) / 2

		// если число из середины равно нашему таргету - одаем индекс
		if nums[mid] == target {
			return mid
		}

		// если взятое число из середины больше таргета -  откидываем правую часть
		if nums[mid] > target {
			right = mid - 1
			// если взятое число из середины меньше таргета -  откидываем левую часть
		} else {
			left = mid + 1
		}
	}

	// если ничего не нашли возвращаем -1
	return -1
}

func main() {

	arr := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// вернет  идекс 2
	fmt.Println(binarySearch(arr, 12))

	// верет -1
	fmt.Println(binarySearch(arr, 333))
}
