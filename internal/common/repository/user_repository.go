package repository

import (
	"database/sql"
	"shortener-auth/internal/common/model"
)

type UserRepository interface {
	GetUserByLogin(login string) (*model.User, error)
	CreateUser(login, password string) error
}

type UserRepositoryPostgresql struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) *UserRepositoryPostgresql {
	return &UserRepositoryPostgresql{conn: conn}
}

func (r *UserRepositoryPostgresql) getNextId() (int, error) {
	result := r.conn.QueryRow("SELECT nextval(pg_get_serial_sequence('user', 'id'))")

	var nextId int

	err := result.Scan(&nextId)
	if err != nil {
		return 0, err
	}

	return nextId, nil
}

func (r *UserRepositoryPostgresql) CreateUser(login, password string) error {
	id, err := r.getNextId()

	if err != nil {
		return err
	}
	role := "ROLE_USER"
	_, err = r.conn.Exec(`
		INSERT INTO "user" (id, login, password, role) VALUES ($1, $2, $3, $4);
		`,
		id, login, password, role,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryPostgresql) GetUserByLogin(login string) (*model.User, error) {
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
