package creational

import (
	"errors"
	"fmt"
)

//PaymentMethod defines a way of paying in the shop. This factory method returns
//objects that implements this interface
type PaymentMethod interface {
	Pay(amount float32) string
}

//Our current implemented Payment methods are described here
const (
	Cash      = 1
	DebitCard = 2
)

//CreatePaymentMethod returns a pointer to a PaymentMethod object or an error
//if the method is not registered. We used "new" operator to return the pointer
//but we could also used &Type{} althought new makes it more readable for
//newcomers could be confusing
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(NewDebitCardPM), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f payed using cash\n", amount)
}

func (c *DebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using debit card\n", amount)
}

type NewDebitCardPM struct{}

func (d *NewDebitCardPM) Pay(amount float32) string {
	return fmt.Sprintf("%#0.2f payed using debit card (new)\n", amount)
}
