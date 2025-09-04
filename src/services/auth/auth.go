package authServices

import (
	"database/sql"

	"github.com/programadorisgod/auth-service/src/config"
	"github.com/programadorisgod/auth-service/src/models/user"
)

func SaveUser(u *user.UserRegister) (int, error) {
	var id int
	err := config.DB.QueryRow(
		"INSERT INTO USERS (email, name, pass) VALUES ($1, $2, $3) RETURNING id",
		u.Email, u.Name, u.Pass,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func FindUser(email string) (*user.User, error) {

	var u user.User

	err := config.DB.QueryRow(
		"SELECT id, name, email, pass, create_at FROM USERS WHERE email = $1",
		email,
	).Scan(&u.Id, &u.Name, &u.Email, &u.Pass, &u.Create_at)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &u, nil

}
