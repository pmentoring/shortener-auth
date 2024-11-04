package repository

import "database/sql"

type RegisterRepository interface {
	CreateUser(login, password string) error
}

type RegisterRepositoryPostgresql struct {
	conn *sql.DB
}

func NewRegisterRepository(conn *sql.DB) *RegisterRepositoryPostgresql {
	return &RegisterRepositoryPostgresql{conn: conn}
}

func (repo RegisterRepositoryPostgresql) getNextId() (int, error) {
	result := repo.conn.QueryRow("SELECT nextval(pg_get_serial_sequence('user', 'id'))")

	var nextId int

	err := result.Scan(&nextId)
	if err != nil {
		return 0, err
	}

	return nextId, nil
}

func (repo RegisterRepositoryPostgresql) CreateUser(login, password string) error {
	id, err := repo.getNextId()

	if err != nil {
		return err
	}
	role := "ROLE_USER"
	_, err = repo.conn.Exec(`
		INSERT INTO "user" (id, login, password, role) VALUES ($1, $2, $3, $4);
		`,
		id, login, password, role,
	)

	if err != nil {
		return err
	}

	return nil
}
