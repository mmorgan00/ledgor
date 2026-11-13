package store

import (
	"context"
	"github.com/jackc/pgx/v5"
)

type Store struct {
	conn *pgx.Conn
}

func New(ctx context.Context, url string) (*Store, error) {
	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		return nil, err
	}
	return &Store{conn: conn}, nil
}

func Connect(cxnStr string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), cxnStr)

	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *Store) Close(ctx context.Context) error {
	return s.conn.Close(ctx)
}
