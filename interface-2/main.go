package main

import (
	"fmt"

	"github.com/shubhamrasal/go-practice/interface-2/pkg/payments"
	"github.com/shubhamrasal/go-practice/interface-2/pkg/payments/creditcard"
	"github.com/shubhamrasal/go-practice/interface-2/pkg/payments/upi"
)

func Checkout(method payments.PaymentMethod, amount float64) string {

	msg := method.Pay(amount)
	return msg
}

func main() {

	fmt.Println("payment interface example")

	shubhamUPI := upi.UPIPayment{UpiID: "shubham@okicici", App: "Gpay"}

	msg := Checkout(shubhamUPI, 24.65)

	fmt.Printf("Payment successful: %s", msg)

	shubhamCard := creditcard.CreditCardPayment{CardNumber: "1234567890123456", ExpiryDate: "01/2026", CVV: "123"}

	msg = Checkout(shubhamCard, 24.65)

	fmt.Printf("Payment successful: %s", msg)
}
