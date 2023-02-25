package application

import (
	"time"
	"wallet/internal/domain/wallet"
)

type WalletApplication struct {
	WalletRepository wallet.WalletRepository
}

func NewWalletApplication(walletRepository wallet.WalletRepository) WalletApplication {
	return WalletApplication{
		WalletRepository: walletRepository,
	}
}

func (a WalletApplication) Add(userID, amount int64) (int64, error) {
	userIDValue, err := wallet.NewUserID(userID)
	if err != nil {
		return 0, err
	}

	amountValue, err := wallet.NewAmount(amount)
	if err != nil {
		return 0, err
	}

	balance, err := a.WalletRepository.GetBalance(userIDValue)
	if err != nil {
		return 0, err
	}

	// TODO: Add SQL Transaction
	if balance-amount < 0 {
		return 0, wallet.InsufficientBalanceError
	}

	walletChange := wallet.NewWalletChange(userIDValue, amountValue)

	referenceID, err := a.WalletRepository.Add(walletChange)
	if err != nil {
		return 0, err
	}

	return referenceID, nil
}

func (a WalletApplication) GetBalance(userID int64) (int64, error) {
	userIDValue, err := wallet.NewUserID(userID)
	if err != nil {
		return 0, err
	}

	return a.WalletRepository.GetBalance(userIDValue)
}

func (a WalletApplication) CalculateTotalAmount() (int64, error) {
	return a.WalletRepository.GetTotalAmountAfter(time.Now().Add(-1 * 24 * time.Hour))
}
