package main

import "fmt"

type Human struct {
	Name string
	Age  uint8
	City string
}

func (h *Human) setName(name string) {
	h.Name = name
}

func (h *Human) setAge(age uint8) {
	h.Age = age
}

func (h *Human) setCity(city string) {
	h.City = city
}

// ShowInfo Родительский метод
func (h *Human) ShowInfo() {
	fmt.Printf("[Human] Name: %s, Age: %d, City: %s.\n", h.Name, h.Age, h.City)
}

type Action struct {
	//встраиваем родительский класс (аналог наследования)
	Human
}

// ShowInfo Переопределенный родительский метод
func (a *Action) ShowInfo() {
	fmt.Printf("[Action] Name: %s, Age: %d, City: %s.\n", a.Name, a.Age, a.City)
}

func (a *Action) drink(liquid string) {
	fmt.Printf("%s drink %s\n", a.Name, liquid)
}

// NewAction делаем функцию коструктор
func NewAction(name string, age uint8, city string) *Action {
	a := Action{}

	// Вызываем родительские методы Human
	a.setName(name)
	a.setAge(age)
	a.setCity(city)

	return &a
}

func main() {
	action := NewAction("Vitaliy", 23, "Moscow")

	// Вызываем переопределнный метод
	action.ShowInfo()

	// Также мы имеем возможность обратиться к родительскому методу
	action.Human.ShowInfo()

}
