package persistance

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
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
	res, err := w.DB.Exec("INSERT INTO wallet_changes (user_id, amount, created_at) VALUES(?, ?, ?)",
		change.UserID,
		change.Amount,
		change.CreatedAt.Format("2006-01-02 15:04:05"),
	)
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

func (w walletMySQLRepository) GetTotalAmountAfter(after time.Time) (int64, error) {
	amount := struct {
		TotalAmount float64 `db:"total_amount"`
	}{}

	err := w.DB.Get(&amount, "SELECT SUM(amount) total_amount FROM wallet_changes WHERE ? < created_at;", after.Format("2006-01-02 15:04:05"))
	if err != nil {
		return 0, err
	}

	return int64(amount.TotalAmount), nil
}
