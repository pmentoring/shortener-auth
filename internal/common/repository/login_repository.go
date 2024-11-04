package repository

import (
	"database/sql"
	"shortener-auth/internal/common/model"
)

type LoginRepository interface {
	GetUserByLogin(login string) (*model.User, error)
}

type LoginRepositoryPostgresql struct {
	conn *sql.DB
}

func NewLoginRepository(conn *sql.DB) *LoginRepositoryPostgresql {
	return &LoginRepositoryPostgresql{conn: conn}
}

func (r LoginRepositoryPostgresql) GetUserByLogin(login string) (*model.User, error) {
	row := r.conn.QueryRow(`SELECT * FROM public.user WHERE login = $1`, login)

	err := row.Err()
	if err != nil {
		return nil, err
	}

	user := new(model.User)

	err = row.Scan(&user.Id, &user.Login, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
