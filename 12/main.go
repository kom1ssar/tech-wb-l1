package main

import "fmt"

// множество строк
func intersectStrings(arr []string) []string {
	// создаем мапку
	hash := make(map[string]struct{})
	var res []string

	// ебжим по массиву
	for _, v := range arr {

		// если есть в мапе - пропускаем
		if _, ok := hash[v]; ok {
			continue
		}

		// добавляем в мапу
		hash[v] = struct{}{}
		// пушим в массив результатов
		res = append(res, v)
	}

	return res
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	res := intersectStrings(arr)

	fmt.Println(res)
}
