package main

import (
	"wallet/internal/application"
	"wallet/internal/infrastructure/persistance"
	"wallet/internal/interface/http"
)

func main() {
	walletRepository := persistance.NewWalletMySQLRepository()
	walletApplication := application.NewWalletApplication(walletRepository)
	http.StartHTTP(walletApplication)
}
