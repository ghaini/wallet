package persistance

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"wallet/internal/domain/wallet"

	_ "github.com/go-sql-driver/mysql"
)

type walletMySQLRepository struct {
	DB *sqlx.DB
}

func NewWalletMySQLRepository(host, user, password, dbname string) (wallet.WalletRepository, error) {
	conn, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname))
	if err != nil {
		panic(err)
	}

	return walletMySQLRepository{DB: conn}, nil
}

func (w walletMySQLRepository) Add(change wallet.WalletChange) (int64, error) {
	res, err := w.DB.Exec("INSERT INTO wallet_changes (user_id, amount) VALUES(?, ?)", change.UserID, change.Amount)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (w walletMySQLRepository) GetBalance(id wallet.UserID) (int64, error) {
	amount := struct {
		Amount float64 `db:"amount"`
	}{}

	err := w.DB.Get(&amount, "SELECT SUM(amount) amount FROM wallet_changes GROUP BY user_id;")
	if err != nil {
		return 0, err
	}

	return int64(amount.Amount), nil
}
