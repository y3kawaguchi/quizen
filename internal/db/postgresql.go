package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	// github.com/lib/pq ...
	_ "github.com/lib/pq"
)

const postgreSQLDriverName = "postgres"

// PostgreSQLConfig ...
type PostgreSQLConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SslMode  string
}

// PostgreSQLConnection ...
type PostgreSQLConnection struct {
	DB *sql.DB
}

// GetDB ...
func (c *PostgreSQLConnection) GetDB() *sql.DB {
	return c.DB
}

// Close ...
func (c *PostgreSQLConnection) Close() error {
	return c.DB.Close()
}

// GetPostgreSQLConfigFromEnv ...
func GetPostgreSQLConfigFromEnv() (*PostgreSQLConfig, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	return &PostgreSQLConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}, nil
}

// ConnectPostgreSQL ...
func ConnectPostgreSQL(c *PostgreSQLConfig) (Connection, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", c.Host, c.Port, c.User, c.Password, c.DBName, c.SslMode)
	db, err := sql.Open(postgreSQLDriverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	connection := &PostgreSQLConnection{
		DB: db,
	}

	return connection, nil
}
