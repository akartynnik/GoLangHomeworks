package main

import (
	"fmt"

	"./person"
)

func main() {
	p := person.NewPerson(2, "Vano", "secret222")
	fmt.Println(p)

	// fmt.Println(p.secret)  // не сработает

	fmt.Println(person.GetSecret(p))
}
