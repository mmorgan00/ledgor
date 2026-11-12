package store

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func Connect(cxnStr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), cxnStr)

	if err != nil {
		return nil, err
	}
	return conn, nil
}
