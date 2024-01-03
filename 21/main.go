package main

import "fmt"

type Cat struct {
}

// добавляем метод DrinkMilk
func (c *Cat) DrinkMilk() {
	fmt.Println("drink milk")
}

type Dog struct {
	Sleep bool
}

// Собака пойдет пить только когда не спит
func (d *Dog) DrinkWatter() {
	if d.Sleep {
		fmt.Println("Dog is sleeping")
		return
	}
	fmt.Println("drink watter")
}

// создаем общий интерфейс с Drink()
type AnimalAdapter interface {
	Drink()
}

// Создаем адаптер собаки
type DogAdapter struct {
	*Dog
}

// реализуем метод интерфейса
// в контексте адаптера нам будет все ранво спит собака или нет
func (d *DogAdapter) Drink() {
	f := d.Dog.Sleep
	d.Dog.Sleep = false

	d.Dog.DrinkWatter()

	d.Dog.Sleep = f
}

// функция конструктор адаптера собаки возвращающая интерфейс
func NewDogAdapter(dog *Dog) AnimalAdapter {
	return &DogAdapter{dog}
}

// создаем структуру адпатера кошки
type CatAdapter struct {
	*Cat
}

// реализуем метод интерфейса
func (c *CatAdapter) Drink() {
	c.Cat.DrinkMilk()
}

// функция конструктор адаптера кошки возвращающая интерфейс
func NewCatAdapter(cat *Cat) AnimalAdapter {
	return &CatAdapter{cat}
}

func main() {
	// создаем адаптер кошки
	catAdapter := NewCatAdapter(&Cat{})
	// создае адаптер собаки
	dogAdapter := NewDogAdapter(&Dog{})

	// вызываем методы адаптера кошки и собаки
	dogAdapter.Drink()
	catAdapter.Drink()

}
