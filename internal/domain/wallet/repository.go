package wallet

type WalletRepository interface {
	Add(change WalletChange) (int64, error)
	GetBalance(id UserID) (int64, error)
}
