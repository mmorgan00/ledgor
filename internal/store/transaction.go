package store

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type TransactionState string

const (
	StatePending TransactionState = "pending"
	StateSettled TransactionState = "settled"
	StateFailed  TransactionState = "failed"
)

type Transaction struct {
	Id             uuid.UUID        `json:"id"`
	Description    *string          `json:"description"`
	IdempotencyKey string           `json:"idempotencyKey"`
	Amount         int              `json:"amount"`
	State          TransactionState `json:"transactionState"`
}

func NewTransaction(amount int, description *string, idemp_key string) Transaction {
	id := uuid.New()
	txn := Transaction{Id: id, IdempotencyKey: idemp_key, Amount: 10, Description: description}
	txn.fill_defaults()
	return txn
}

func (txn *Transaction) fill_defaults() {
	if txn.State == "" {
		txn.State = StatePending
	}
}

func (s *Store) TransactionInsert(t *Transaction) error {
	fmt.Println("Inserting new txn:", t)
	return pgx.BeginFunc(context.Background(), s.conn, func(tx pgx.Tx) error {
		_, err := tx.Exec(context.Background(), `
			INSERT INTO "Transaction" (id, description, idempotencyKey, amount, state)
			VALUES ($1, $2, $3, $4, $5)`,
			t.Id, t.Description, t.IdempotencyKey, t.Amount, t.State)
		return err
	})
}

func (s *Store) TransactionGetById(id uuid.UUID) (*Transaction, error) {
	var txn Transaction
	err := s.conn.QueryRow(context.Background(),
		`SELECT id, description, idempotencyKey, amount, state FROM "Transaction" WHERE id = $1`,
		id).Scan(&txn.Id, &txn.Description, &txn.IdempotencyKey, &txn.Amount, &txn.State)
	if err != nil {
		return nil, err
	}
	return &txn, nil
}
