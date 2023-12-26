package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// выход по таймауту контекста
func exitGoroutineByCtxTimer() {

	// добавляем WaitGroup, чтобы дождаться завершения горнутины
	wg := sync.WaitGroup{}

	// создаем контектс с таймаутом 2 секунды
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	wg.Add(1)
	// создаем горутину
	go func() {
		defer wg.Done()
		for {
			select {
			// если контекст отменился ( в нашем случае мы ждем завершения таймера) - выхдим
			case <-ctx.Done():
				fmt.Println("Exit with context timeout")
				return
			default:
				// выводим рандомное число каждые 100ms
				fmt.Println(rand.Int())
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	// ждем завершеня горутины по таймауту
	wg.Wait()

}
