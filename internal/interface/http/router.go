package http

import (
	"github.com/gin-gonic/gin"
	"wallet/internal/application"
)

func StartHTTP(walletApplication application.WalletApplication) {
	gin.SetMode(gin.ReleaseMode)
	g := gin.Default()
	handler := NewHandler(walletApplication)
	router(g, handler)
	g.Run(":80")
}

func router(engine *gin.Engine, handler Handler) {
	engine.GET("/ping", handler.Pong)

	engine.GET("/user/:user_id/get-balance", handler.GetBalance)
	engine.POST("/add-money", handler.AddMoney)
}
