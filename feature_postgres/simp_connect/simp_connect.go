package simpconnect

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CheckConnection() (*pgx.Conn, error) {
	res, err := pgx.Connect(context.Background(), "postgres://postgres:mol@localhost:5432/postgres")
	if err != nil {
		panic(err)
	}

	err = res.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	return res, nil
}
