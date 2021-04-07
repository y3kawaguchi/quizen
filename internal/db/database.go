package db

import (
	"database/sql"
)

// Connection ...
type Connection interface {
	GetDB() *sql.DB
	Close() error
}
