package persistance

import (
	"wallet/internal/domain/wallet"
)

type walletMySQLRepository struct{}

func NewWalletMySQLRepository() wallet.WalletRepository {
	return walletMySQLRepository{}
}

func (w walletMySQLRepository) Add(change wallet.WalletChange) (int64, error) {
	return 0, nil
}

func (w walletMySQLRepository) GetBalance(id wallet.UserID) (int64, error) {
	return 0, nil
}
