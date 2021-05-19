package pointers_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// ErrorInsufficientFunds Глобальная переменная в рамках пакета
var ErrorInsufficientFunds = errors.New("Нельзя снимать денег больше чем имеется!")

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in method Deposit()? %v\n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrorInsufficientFunds
	}
	w.balance -= amount

	return nil
}

//--
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
