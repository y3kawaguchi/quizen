package db

import (
	"github.com/jmoiron/sqlx"
)

// Connection ...
type Connection interface {
	GetDB() *sqlx.DB
	Close() error
}
