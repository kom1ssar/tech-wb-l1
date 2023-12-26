package main

import (
	"fmt"
	"sync"
)

// завершение с помощью сообщения в канал
func exitGoroutineByChExit() {

	// добавляем WaitGroup, чтобы дождаться завершения горнутины
	wg := sync.WaitGroup{}

	// канал для предачи данных
	ch := make(chan int)
	// канал для выхода
	exit := make(chan interface{}, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			// если что то пришло - завершаем работу горутины
			case <-exit:
				fmt.Println("Exit by message on 'exit' channel ")
				return
			// читаем с канала данные
			case value := <-ch:
				fmt.Println(value)
			}
		}
	}()

	// пишем в канал
	for i := 0; i <= 20; i++ {
		ch <- i
	}

	// закрываем канал тем самым завершаем выполненние горутины
	close(exit)

	// ожидаем завершения горутины
	wg.Wait()
}
