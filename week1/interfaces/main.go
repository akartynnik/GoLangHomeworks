package main

import (
	"fmt"
)

type Payer interface {
	Pay(int) error
}

// ----------------

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает налички")
	}
	w.Cash -= amount
	return nil
}

// ----------------

type Card struct {
	Balance   int
	AccountId int
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Не хватает бабок на карте")
	}
	c.Balance -= amount
	return nil
}

// ----------------

type GooglePay struct {
	Money    int
	GoogleId int
}

func (gp *GooglePay) Pay(amount int) error {
	if gp.Money < amount {
		return fmt.Errorf("Не хватает бабок гугло-аккаунте")
	}
	gp.Money -= amount
	return nil
}

// ----------------

func Buy(p Payer) {
	switch p.(type) {
	case *Wallet:
		fmt.Println("Оплата кэшем")
	case *Card:
		fmt.Println("Карточка")
	case *GooglePay:
		fmt.Println("Гугло-оплата")
	default:
		fmt.Println("Хз что это...")
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка оплаты через %T: %v\n\n", p, err)
		return
	}

	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var payer Payer
	payer = &Card{Balance: 100}
	Buy(payer)

	payer = &GooglePay{Money: 6}
	Buy(payer)
}
