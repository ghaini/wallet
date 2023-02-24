package wallet

type UserID int64

func NewUserID(userID int64) (UserID, error) {
	if userID <= 0 {
		return 0, InvalidUserIDError
	}

	return UserID(userID), nil
}