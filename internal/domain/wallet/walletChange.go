package wallet

import "time"

type WalletChange struct {
	ID        int64
	UserID    UserID
	Amount    Amount
	CreatedAt time.Time
}

func NewWalletChange(userID UserID, amount Amount) WalletChange {
	return WalletChange{
		UserID:    userID,
		Amount:    amount,
		CreatedAt: time.Now().UTC(),
	}
}
