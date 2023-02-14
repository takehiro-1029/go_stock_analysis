package infra

import (
	"context"
	"database/sql"
	"log"
)

type transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) *transaction {
	return &transaction{db: db}
}

func (t *transaction) Transaction(ctx context.Context, f func(ctx context.Context, tx *sql.Tx) error) error {
	// trasaction start
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		// panic
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				log.Printf("failed to MySQL Rollback: %v", err)
			}
			// re-throw panic after Rollback
			panic(p)
		}
		// error
		if err != nil {
			if err = tx.Rollback(); err != nil {
				log.Printf("failed to MySQL Rollback: %v", err)
			}
			return
		}
		// 正常
		if err := tx.Commit(); err != nil {
			log.Printf("failed to MySQL Commit: %v", err)
		}

	}()

	// 主処理実行
	if err := f(ctx, tx); err != nil {
		return err
	}
	return nil
}
