package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	amount Bitcoin
}

// Return amount of the copy of our wallet
// Can be changed to *Wallet to recieve a amount from original wallet but no difference now
func (wallet Wallet) Balance() Bitcoin {
	return wallet.amount
}

func (wallet *Wallet) Deposit(amount Bitcoin) string {
	wallet.amount = wallet.amount + amount
	return fmt.Sprintf("%.f ammount was deposited to wallet %#v", amount, wallet)
}

func (wallet *Wallet) Withdraw(amount Bitcoin) (msg string, e error) {
	if wallet.amount < amount {
		e = errors.New("Withdraw ammount cannot be bigger than actual balance")
	} else {
		msg = fmt.Sprintf("Withdrawed %T %.f from balance", amount, amount)
		wallet.amount -= amount
	}
	return
}
