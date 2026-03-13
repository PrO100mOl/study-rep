package simpsql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, c *pgx.Conn) ([]Task, error) {
	SQLSelect := `
	SELECT id, title, description, completed, created_at, completed_at 
	FROM tasks

	`
	res, err := c.Query(ctx, SQLSelect)
	if err != nil {
		return nil, err
	}
	defer res.Close()
	templist := []Task{}
	for res.Next() {
		var temp Task
		err := res.Scan(
			&temp.ID,
			&temp.Title,
			&temp.Description,
			&temp.Completed,
			&temp.Created_at,
			&temp.Completed_at,
		)
		if err != nil {
			return nil, err
		}
		templist = append(templist, temp)
	}
	return templist, err
}
