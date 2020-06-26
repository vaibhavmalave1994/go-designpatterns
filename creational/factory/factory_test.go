package factory

import (
	"strings"
	"testing"
)

func TestCashPaymentMethod(t *testing.T) {
	var payment PaymentMethod
	var err error
	payment, err = GetPaymentMethod(Cash)

	if err != nil {
		t.Error("Cash payment method must be implemented.")
	}

	result := payment.Pay(10.2)
	if !strings.Contains(result, "paid using cash payment method") {
		t.Error("The cash payment method result wasn't correct")
	}
	t.Log("result", result)
}

func TestDebitCardPaymentMethod(t *testing.T) {
	var payment PaymentMethod
	var err error
	payment, err = GetPaymentMethod(DebitCard)

	if err != nil {
		t.Error("Debit card payment method must be implemented.")
	}

	result := payment.Pay(10.2)
	if !strings.Contains(result, "paid using debit card payment method") {
		t.Error("The debit card payment method result wasn't correct")
	}
	t.Log("result", result)
}

func TestPaymentMethodDoesNotExist(t *testing.T) {
	var err error
	paymentMethod := 10
	_, err = GetPaymentMethod(paymentMethod)

	if err == nil {
		t.Errorf("Payment method with id %d must return an error.", paymentMethod)
	}
}
