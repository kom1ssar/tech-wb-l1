package main

import "fmt"

// создаем конвейр чисел
func pipelineNumbs(nums []int) {

	// создаем 2 канала, я решил использовать не буферизированные
	var (
		firstCh  = make(chan int)
		secondCh = make(chan int)
	)

	// создаем горутину
	go func() {
		// закрываем канал по завершению горутины
		defer close(secondCh)

		// читем из первого канала
		for num := range firstCh {
			// пишем резульат n * 2 в второй канал
			secondCh <- num * 2
		}
	}()
	// создаем горутину
	go func() {
		// закрываем канал по завершению горутины
		defer close(firstCh)

		// бежим по массиву
		for _, num := range nums {
			// пишем значения из массива в первый канал
			firstCh <- num
		}

	}()

	// читаем из второго канала
	for num := range secondCh {
		// выводим результат в stdout
		fmt.Println(num)
	}

}

func main() {

	nums := []int{13, 11, 4, 7, 5, 6, 7, 25, 22, 10}
	pipelineNumbs(nums)
}
