package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Решение через Atomic
func sequenceByAtomic(nums []int64) int64 {
	var (
		wg     sync.WaitGroup
		result int64
	)

	for _, v := range nums {
		wg.Add(1)
		go func(num int64) {
			defer wg.Done()
			// Атомик нам гарантирует потокобезопасное выполнение
			atomic.AddInt64(&result, num*num)
		}(v)
	}

	wg.Wait()

	return result
}

// Решение через Mutex
func sequenceByMutex(nums []int64) int64 {
	var (
		wg     sync.WaitGroup
		mu     sync.Mutex
		result int64
	)

	for _, v := range nums {
		wg.Add(1)
		go func(num int64) {
			// снимаем блокировку
			defer func() {
				mu.Unlock()
				wg.Done()
			}()
			// блокируемся для cинхронизации доступа к переменной
			mu.Lock()

			result += num * num

		}(v)
	}

	wg.Wait()

	return result

}

func main() {
	nums := []int64{2, 4, 6, 8, 10}

	fmt.Println(sequenceByAtomic(nums))
	fmt.Println(sequenceByMutex(nums))
}
