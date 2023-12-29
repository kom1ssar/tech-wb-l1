package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// создаем структуру каунтера
type Counter struct {
	// для такой простой операции как инкремент, нам отлично подойдет атомик. Мьютекс тут точо будет излишнем)
	accumulate atomic.Uint64
}

// потокобезопасно инкрементируем
func (c *Counter) Increment() {
	// безопасно добавляем единицу к нашему  accumulate
	c.accumulate.Add(1)
}

// получаем значение атомика
func (c *Counter) Get() uint64 {
	return c.accumulate.Load()
}

// функция конструктор
func NewCounter() *Counter {
	return &Counter{}
}

func main() {

	// восопльзуемся WaitGroup для ожидания завершения всех горутин
	wg := sync.WaitGroup{}

	// создаем нашу структуру
	c := NewCounter()

	// лимит горутин
	limit := 1000

	for i := 0; i <= limit; i++ {

		// создаем горутины
		wg.Add(1)
		go func() {
			// по заврешению горутины сообщаем в wg
			defer wg.Done()
			// инкрементимся
			c.Increment()
		}()

	}

	// дожидаемся завершения всех горунтин
	wg.Wait()
	//выводим результат
	fmt.Println(c.Get())

}
