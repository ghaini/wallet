package main

import (
	"log"
	"wallet/config"
	"wallet/internal/application"
	"wallet/internal/infrastructure/persistance"
	"wallet/internal/interface/cron"
	"wallet/internal/interface/http"
)

func main() {
	cfg := config.Init()
	walletRepository, err := persistance.NewWalletMySQLRepository(cfg.MySQL.Host, cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.DBName)
	if err != nil {
		log.Fatalln(err)
	}

	walletApplication := application.NewWalletApplication(walletRepository)
	go cron.StartCron(walletApplication)
	http.StartHTTP(walletApplication)
}
