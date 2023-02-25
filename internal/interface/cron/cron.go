package cron

import (
	"github.com/jasonlvhit/gocron"
	"log"
	"wallet/internal/application"
)

type Cron struct {
	walletApplication application.WalletApplication
}

func StartCron(walletApplication application.WalletApplication) {
	c := Cron{walletApplication: walletApplication}
	gocron.Every(1).Day().At("00:00").Do(c.CalculateTotalAmount)
	<-gocron.Start()
}

func (c Cron) CalculateTotalAmount() {
	totalAmount, err := c.walletApplication.CalculateTotalAmount()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Total Amount of all yesterday's transactions: ", totalAmount)
}
