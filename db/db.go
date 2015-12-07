package db

import (
	"database/sql"
	"fmt"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

type DBProxy struct {
	Db *sql.DB
}

func (dbProxy *DBProxy) Init(config *Config) error {
	var err error
	address := fmt.Sprintf("%s:%s", config.DBHost, config.DBPort)
	connStr := fmt.Sprintf("%s:%s@/%s?default-character-set=utf8", config.username, config.password, address)
	Db, err = sql.Open("mysql", connStr)
	if err != nil {
		return err
	}
	Db.SetMaxIdleConns(config.DBIdleConns)
	Db.SetMaxOpenConns(config.DBOpenConns)

	return nil
}
