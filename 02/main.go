package main

import (
	"fmt"
	"sync"
)

// задаем колличество воркеров
const workersLimit = 2

// решение через воркеров
func squareWorkerPool(nums []int) {
	wg := &sync.WaitGroup{}

	// создаем 2 канала. Первый для передачи значение в воркер, второй для получения результата выполнения
	toProcessed, processed := make(chan int, workersLimit), make(chan int, workersLimit)

	// создаем нужное кол-во воркеров заданных константой  `workersLimit`
	for i := 0; i <= workersLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// передаем созданные каналы в воркер
			squareWorker(toProcessed, processed)
		}()
	}

	go func() {
		// нам обязательно нужно позаботиться о закрытие канала
		defer close(toProcessed)

		// пишем значения из нашего массива в канал
		for _, v := range nums {
			toProcessed <- v
		}

	}()

	go func() {
		wg.Wait()
		// чтобы не повиснуть создаем горутину, которая закроет наш канал, когда все воркеры закончат работу
		close(processed)
	}()

	// читаем результат с канала
	for res := range processed {
		fmt.Println(res)
	}

}

// решение через воркеров
func squareWorker(toProcessed <-chan int, processed chan<- int) {
	for {
		select {
		// получаем значение из канала
		case value, ok := <-toProcessed:
			if !ok {
				return
			}

			// пишем результат в канал
			processed <- value * value
		}
	}
}

// Создаем горутину на каждое число
func square(nums []int) {
	wg := &sync.WaitGroup{}

	for _, v := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(v)
	}

	wg.Wait()
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	squareWorkerPool(nums)
	square(nums)

}
