package simpconnect

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func CheckConnection(ctx context.Context) (*pgx.Conn, error) {
	gg := os.Getenv("abiba")
	return pgx.Connect(ctx, gg)
}
