package simpsql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func DeleteRow(ctx context.Context, c *pgx.Conn) error {
	SQLDelete := `
		DELETE FROM tasks
		WHERE id = 4;
	`
	_, err := c.Exec(ctx, SQLDelete)
	return err
}
