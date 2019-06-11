package main

import (
	"fmt"
)

type Person struct {
	Id        int
	FirstName string
	LastName  string
}

func (p *Person) SetFullName(firstName, lastName string) {
	p.FirstName = firstName
	p.LastName = lastName
}

func main() {
	person := Person{1, "Лекс", "АйТиБорода"} // быстрое создание и инициализация структуры
	person.SetFullName("Женя", "Соер")
	fmt.Println(person) // {1 Женя Соер}
}
