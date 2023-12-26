package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// выходим по закрытию контекста
func exitGoroutineByCtx() {

	// добавляем WaitGroup, чтобы дождаться завершения горнутины
	wg := sync.WaitGroup{}

	// создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)

	// запускаем горутину
	go func() {
		defer wg.Done()

		for {
			select {
			// если контекст отменен - выходим
			case <-ctx.Done():
				fmt.Println("Exit by ctx closed")
				return
			default:

				// выводим в консоль рандомное число каждые 100ms
				fmt.Println(rand.Int())
				time.Sleep(time.Millisecond * 100)
			}

		}
	}()

	// дадим поработать горутине пару секунд
	time.Sleep(time.Second * 2)

	// завершаем контекст
	cancel()

	// дождиаемся завершения горутины
	wg.Wait()
}
