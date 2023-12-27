package main

import "fmt"

func intersect(first []int, second []int) []int {
	// делаем мапку
	hash := make(map[int]struct{})
	var res []int

	//  бежим по первому массиву
	for _, v := range first {

		// если уже есть мапе - пропускаем
		if _, ok := hash[v]; ok {
			continue
		}

		// добавляем в мапу
		hash[v] = struct{}{}
		// пушим в результат
		res = append(res, v)
	}

	//  бежим по второму массиву
	for _, v := range second {

		// если уже есть мапе - пропускаем
		if _, ok := hash[v]; ok {
			continue
		}
		// добавляем в мапу
		hash[v] = struct{}{}
		// пушим в результат
		res = append(res, v)
	}

	return res

}

func main() {
	first := []int{1, 3, 4, 4, 5, 6, 8}
	second := []int{2, 3, 5, 5, 1, 9, 8, 10}

	res := intersect(first, second)
	fmt.Println(res)
}
