package registry

import (
	"database/sql"
)

type Registry interface {
	DB() *sql.DB
}

type registry struct {
	db *sql.DB
}

// NewRegistry 新しいDIコンテナを作成する
func NewRegistry(db *sql.DB) Registry {
	return &registry{
		db: db,
	}
}

func (r *registry) DB() *sql.DB {
	return r.db
}
