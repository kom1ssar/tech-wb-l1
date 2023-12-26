package main

import (
	"fmt"
	"sync"
)

// для удобства создадим структуру
type SyncMap struct {
	// добавим RWMutex для синхронизации.
	// Используем его для повышения производительности - достижением паралельного чтения
	rm   *sync.RWMutex
	data map[int]int
}

// метод добавления
func (s *SyncMap) Set(key int, value int) {
	// разблокировка по завершению
	defer s.rm.Unlock()

	// блокируем
	s.rm.Lock()
	s.data[key] = value
}

// метод чтения
func (s *SyncMap) Get(key int) (int, bool) {
	// разблокировка по завершению
	defer s.rm.RUnlock()

	// блокируем
	s.rm.RLock()
	v, ok := s.data[key]
	return v, ok
}

// фуункция конструктор
func NewSyncMap() *SyncMap {
	return &SyncMap{
		rm:   &sync.RWMutex{},
		data: make(map[int]int),
	}

}

func main() {

	// добавим WaitGroup для того чтобы дождаться выполнения наших горутин
	wg := sync.WaitGroup{}

	// создаем нашу мапку
	s := NewSyncMap()

	// сделаем цикл на запуск 10 горутин на запись и 10 на чтение
	for i := 0; i <= 10; i++ {

		wg.Add(1)
		// создаем горутину на запись
		go func(num int) {
			defer wg.Done()
			// записываем в нашу мапу
			s.Set(num, num)
		}(i)

		wg.Add(1)
		// создаем горутину на чтение
		go func(num int) {
			defer wg.Done()
			// читаем из мапы
			v, _ := s.Get(num)
			// выводим в  stdout
			fmt.Println(v)
		}(i)
	}

	// дожидаемся завершения наших горутин
	wg.Wait()

}
