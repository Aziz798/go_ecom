package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func NewMysqlStorage(cfg mysql.Config) (*sql.DB, error) {
	dsn := cfg.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
	}
	return db, nil
}
