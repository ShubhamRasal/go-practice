package creditcard

import "fmt"

type CreditCardPayment struct {
	CardNumber string
	ExpiryDate string
	CVV        string
}

func (c CreditCardPayment) Pay(amount float64) string {
	// logic to deduct amount from credit card
	msg := fmt.Sprintf("Credit card payment of %.2f successfully completed using %s \n", amount, c.CardNumber)
	return msg
}
