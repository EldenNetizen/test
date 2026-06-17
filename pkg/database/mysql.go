package database

import (
	"github.com/jmoiron/sqlx"
)

const (
	_root   = "127.0.0.1:3306"
	_dbName = "mysql"
)

type MySQLManager struct {
	db *sqlx.DB
}

func NewMySqlManager() *MySQLManager {
	return &MySQLManager{}
}

func (mySqlManager *MySQLManager) Connect() error {
	var err error
	mySqlManager.db, err = sqlx.Open(_dbName, _root)
	return err
}
