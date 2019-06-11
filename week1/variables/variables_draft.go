package main

import (
	"fmt"
)

func main() {
	var num0 int // без инициализации
	var num1 int = 1
	var num2 = 20                      // компилятор умный, сам доперт
	num := 25                          // а вот эта штука объявляет новуб переменную вообще кратко
	var foo, bar string = "foo", "bar" // объявление с инициализацией
	foo, bar = "да", "ладно"           // присвоение

	fmt.Println(num0, num1, num2, num, foo, bar)
}
