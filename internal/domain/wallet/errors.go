package wallet

import "fmt"

var (
	InvalidUserIDError       = fmt.Errorf("invalid userID")
	InsufficientBalanceError = fmt.Errorf("there is insufficient balance")
)
