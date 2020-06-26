package factory

import (
	"errors"
	"fmt"
)

//constants for payment methods
const (
	Cash      = 1
	DebitCard = 2
)

//all payment methods should implement this interface
type PaymentMethod interface {
	Pay(amout float64) string
}

type CashPaymentMethod struct{}      //payment method for cash
type DebitCardPaymentMethod struct{} //payment method for debit card

func (c *CashPaymentMethod) Pay(amount float64) string {
	return fmt.Sprintf("Payment of %0.2f paid using cash payment method.", amount)
}

func (c *DebitCardPaymentMethod) Pay(amount float64) string {
	return fmt.Sprintf("Payment of %0.2f paid using debit card payment method.", amount)
}

func GetPaymentMethod(paymentMethod int) (PaymentMethod, error) {
	switch paymentMethod {
	case 1:
		return &CashPaymentMethod{}, nil
	case 2:
		return &DebitCardPaymentMethod{}, nil
	default:
		return nil, errors.New("Not implemented yet")
	}
}
