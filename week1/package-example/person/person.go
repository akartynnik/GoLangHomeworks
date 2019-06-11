package person

var (
	Public  = 1 // экспортируемая переменная
	private = 1
)

type Person struct {
	ID     int    // экспортируемое поле
	Name   string // экспортируемое поле
	secret string // приватное поле
}

func (p Person) UpdateSecret(secret string) {
	p.secret = secret
}
