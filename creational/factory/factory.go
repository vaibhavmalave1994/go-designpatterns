package factory

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amout float64) string
}

type Cash struct{}
type DebitCard struct{}

func (c *Cash) Pay(amount float64) string {
	return fmt.Sprintf("Payment of %0.2f paid using cash payment method.")
}

func (c *DebitCard) Pay(amount float64) string {
	return fmt.Sprintf("Payment of %0.2f paid using debit card payment method.")
}

func GetPaymentMethod(paymentMethod int) (PaymentMethod, error) {
	return nil, errors.New("Not implemented yet")
}
