package main

import (
	"fmt"
)

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	var account = Account{
		Id:   1,
		Name: "Alex",
		Person: Person{
			Name: "Volodia",
		},
	}

	account.Owner = Person{1, "Alex", "ItBeard"}
	fmt.Println(account)
	fmt.Println(account.Name)        // выведет Alex
	fmt.Println(account.Person.Name) // выведет Volodia
}
