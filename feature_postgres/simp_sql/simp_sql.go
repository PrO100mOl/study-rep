package simpsql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTable(c *pgx.Conn, ctx context.Context) error {
	SQLQuery := `
	Create Table IF NOT EXISTS tasks(
		id SERIAL PRIMARY KEY,
		title VARCHAR(45) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		completed BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		completed_at TIMESTAMP,

		UNIQUE(title)
	);
	`
	_, err := c.Exec(ctx, SQLQuery)
	if err != nil {
		return err
	}
	return nil
}
