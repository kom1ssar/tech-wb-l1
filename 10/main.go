package main

import "fmt"

// групируем температуру с шагом в 10 градусов
func groupTemp(temps []float64) {
	// создаем мапку
	tempMap := make(map[int][]float64)

	// итерируемся по массиву температур
	for _, temp := range temps {
		// округляем
		t := int(temp) / 10 * 10
		// добавляем в мапку
		tempMap[t] = append(tempMap[t], temp)
	}

	// выводим
	fmt.Println(tempMap)

}

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groupTemp(temp)
}
