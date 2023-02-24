package http

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wallet/internal/application"
)

type Handler struct {
	WalletApplication application.WalletApplication
}

func NewHandler(walletApplication application.WalletApplication) Handler {
	return Handler{WalletApplication: walletApplication}
}

func (h Handler) Pong(ctx *gin.Context) {
	ctx.String(200, "pong")
}

func (h Handler) AddMoney(ctx *gin.Context) {
	req := AddMoneyRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return
	}

	referenceID, err := h.WalletApplication.Add(req.UserID, req.Amount)
	if err != nil {
		ctx.String(400, err.Error())
		return
	}

	ctx.JSON(200, AddMoneyResponse{ReferenceID: referenceID})
}

func (h Handler) GetBalance(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	balance, err := h.WalletApplication.GetBalance(userID)
	if err != nil {
		ctx.String(400, err.Error())
		return
	}

	ctx.JSON(200, GetBalanceResponse{Balance: balance})
}
