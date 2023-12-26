package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// воркер читает с канала и выводит в stdout
func worker(id int, toProcessed <-chan string) {
	for {
		select {
		case value, ok := <-toProcessed:
			if !ok {
				fmt.Printf("\nWorker id: %d stopped", id)
				return
			}
			// выводим сообщение из канала
			fmt.Println(value)

		}

	}
}

func workerPool(workerCount int) {

	var wg sync.WaitGroup

	// канал для передачи сообщеий в воркер
	toProcessedCh := make(chan string)

	// создаем заданное кол-во воркеров
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, toProcessedCh)
		}(i)

	}

	// создаем канал на системный выозов
	stopCh := make(chan os.Signal, 1)

	// попдисываемся на событие завершения
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	// создаим отдельную горунутину которая будет ждать сис-колл на завершение
	go func() {
		//  ожидаем сообщение из канала
		<-stopCh
		// закрываем канал
		close(toProcessedCh)
	}()

	// бесконечно пишем в канал
	go func() {

		for i := 0; ; i++ {
			toProcessedCh <- fmt.Sprintf("message %d", i)
			// сделаем небольшую задержку
			time.Sleep(time.Millisecond * 2)
		}

	}()

	// дожидаемся закрытия всех воркеров
	wg.Wait()

}

func main() {
	var n int
	fmt.Print("Worker limit: ")
	fmt.Scan(&n)

	workerPool(n)
}
