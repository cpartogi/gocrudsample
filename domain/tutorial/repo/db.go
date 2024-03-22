package repo

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// New will
func New(db DBTX) *Queries {
	return &Queries{db: db}
}

// Queries will
type Queries struct {
	db DBTX
}

// WithTx will
func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db: tx,
	}
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

// func NewStore(db *sql.DB) tutorial.TutorialRepoInterface {
// 	return &SQLStore{
// 		db:      db,
// 		Queries: New(db),
// 	}
// }
