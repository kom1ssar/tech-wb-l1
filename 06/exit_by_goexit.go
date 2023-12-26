package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// выходим с помощью Goexit
func exitGoroutineByGoExit() {

	// добавляем WaitGroup, чтобы дождаться завершения горнутины
	wg := sync.WaitGroup{}

	wg.Add(1)
	// создаем горутину
	go func() {
		defer func() {
			fmt.Println("Goroutine stopped by Goexit")
			wg.Done()
		}()

		fmt.Println("Goroutine is running")
		// ждем 200ms после запуска горутины
		time.Sleep(time.Millisecond * 200)

		// завершаем горутину с помощью Goexit
		runtime.Goexit()

		// сюда мы никогда не попадаем
		fmt.Println("Unreached code")
	}()

	wg.Wait()

}
