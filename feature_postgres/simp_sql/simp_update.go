package simpsql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, c *pgx.Conn) error {
	SQLUpdate := `
	UPDATE tasks
	SET completed_at = $1	
	WHERE id = 3
	`
	_, err := c.Exec(ctx, SQLUpdate, time.Now())
	return err
}

func UpdateTask(ctx context.Context, c *pgx.Conn, t Task) error {
	SQLQuery := `
	UPDATE tasks
	SET title = $2, description = $3, completed = $4, created_at = $5, completed_at=$6
	WHERE id = $1;
	`
	_, err := c.Exec(ctx, SQLQuery, t.ID, t.Title, t.Description, t.Completed, t.Created_at, t.Completed_at)
	if err != nil {
		return err
	}
	return nil
}
