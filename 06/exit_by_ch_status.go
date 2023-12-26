package main

import (
	"fmt"
	"sync"
)

// выходи по закрытию канала
func exitGoroutineByChStatus() {

	// добавляем WaitGroup, чтобы дождаться завершения горнутины
	wg := sync.WaitGroup{}
	// канал для предачи данных
	ch := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			value, ok := <-ch
			// если канал закрыт - выходим
			if !ok {
				fmt.Println("Exit by status close")
				return
			}
			fmt.Println(value)
		}
	}()

	// пишем в канал
	for i := 0; i <= 20; i++ {
		ch <- i
	}

	// закрываем канал тем самым завершаем выполненние горутины
	close(ch)

	// ожидаем завершения горутины
	wg.Wait()
}
