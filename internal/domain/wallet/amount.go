package wallet

type Amount int64

func NewAmount(amount int64) (Amount, error) {
	return Amount(amount), nil
}
