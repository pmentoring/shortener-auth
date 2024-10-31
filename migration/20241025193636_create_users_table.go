package migration

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateUsersTable, downCreateUsersTable)
}

func upCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE "user" (
			id SERIAL PRIMARY KEY,
			login VARCHAR(32) NOT NULL,
			password TEXT NOT NULL,
			role VARCHAR(32) NOT NULL,
			created_at date DEFAULT CURRENT_TIMESTAMP
		);

		CREATE INDEX user_login_idx ON "user" (login);
	`)

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		CREATE TABLE user_role (
			code VARCHAR(32) PRIMARY KEY,
			name VARCHAR(64) NOT NULL 
		)
	`)

	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO 
		    user_role (code, name)
		VALUES 
			('ROLE_USER', 'User'),
			('ROLE_ADMIN', 'Admin')
	`)

	return err
}

func downCreateUsersTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		DROP TABLE "user";
		DROP TABLE user_role;
	`)
	return err
}
