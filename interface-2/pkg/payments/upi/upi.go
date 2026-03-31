package upi

import "fmt"

type UPIPayment struct {
	UpiID string
	App   string
}

func (u UPIPayment) Pay(amount float64) string {

	// logic to deduct amount from you upi id
	msg := fmt.Sprintf("UPI payment of %.2f successfully completed using %s\n", amount, u.UpiID)
	return msg
}
