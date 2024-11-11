package database

import (
	"context"

	"api/internal/shared/config"
	e "api/internal/shared/error"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Pool(config config.Config) (*sqlx.DB, func(), error) {
	db, err := sqlx.Open("postgres", config.DatabaseUrl)
	if err != nil {
		return nil, func() {}, err
	}
	return db, func() {
		db.Close()
	}, nil
}

type Tx struct {
	tx  *sqlx.Tx
	ctx context.Context
}

func Begin(db *sqlx.DB) (Tx, error) {
	tx, err := db.Beginx()
	if err != nil {
		return Tx{}, err
	}
	return Tx{tx, context.Background()}, nil
}

func (t *Tx) Rollback() error {
	if t == nil {
		return e.ErrorBuilder(e.Unknown).Build()
	}
	err := t.tx.Rollback()
	t = nil
	return err
}

func (t *Tx) Commit() error {
	if t == nil {
		return e.ErrorBuilder(e.Unknown).Build()
	}
	err := t.tx.Commit()
	t = nil
	return err
}

func (t *Tx) Tx() (*sqlx.Tx, error) {
	if t == nil {
		return nil, e.ErrorBuilder(e.Unknown).Build()
	}
	return t.tx, nil
}

func (t *Tx) Ctx() (context.Context, error) {
	if t == nil {
		return nil, e.ErrorBuilder(e.Unknown).Build()
	}
	return t.ctx, nil
}
