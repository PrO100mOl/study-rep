package simpsql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InsertRow(c *pgx.Conn, ctx context.Context, t Task) error {
	SQLQuery := `
	INSERT INTO tasks (title,description,completed,created_at)
	VALUES ($1, $2, $3, $4);
	`

	_, err := c.Exec(ctx, SQLQuery, t.Title, t.Description, t.Completed, t.Created_at)
	if err != nil {
		return err
	}
	return nil
}
