package wallet

import "time"

type WalletRepository interface {
	Add(change WalletChange) (int64, error)
	GetBalance(id UserID) (int64, error)
	GetTotalAmountAfter(after time.Time) (int64, error)
}
