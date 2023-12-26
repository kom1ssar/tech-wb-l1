package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Завершаем работу не менее чем через N секунд
func consistentlyByTimer(timerSecond uint) {

	// создаем канал
	ch := make(chan int)

	// пишем в канал кажыде полсекунды
	go func() {
		for {
			ch <- rand.Int()
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// запускаем таймер
	timer := time.NewTimer(time.Duration(timerSecond) * time.Second)

	for {
		select {
		// читаем из канала
		case value, ok := <-ch:
			if !ok {
				fmt.Println("channel is closed")
				return
			}

			// выводим сообщение из канала
			fmt.Println(value)

			// проверяем таймер
		case <-timer.C:
			// закрываем канал
			close(ch)
			fmt.Println("Closed by timer")
			// выходим по истечению таймера
			return

		}
	}
}

func main() {
	consistentlyByTimer(5)
}
