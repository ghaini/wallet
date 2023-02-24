package http

type AddMoneyResponse struct {
	ReferenceID int64 `json:"reference_id"`
}

type GetBalanceResponse struct {
	Balance int64 `json:"balance"`
}
