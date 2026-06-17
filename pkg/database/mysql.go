package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	mysqlRoot = "test_user:123456@tcp(127.0.0.1:3306)/test"
	mysqlName = "mysql"
)

type MySQLManager struct {
	db *sqlx.DB
}

func NewMySqlManager() *MySQLManager {
	return &MySQLManager{}
}

func (mySqlManager *MySQLManager) Connect() error {
	var err error
	mySqlManager.db, err = sqlx.Open(mysqlName, mysqlRoot)
	return err
}

func Query[T any](mySQLManager *MySQLManager, obj *T, sql string, params ...string) error {
	err := mySQLManager.db.Get(obj, sql, "1")
	return err
}

func (mySQLManager *MySQLManager) Insert(sql string, params ...any) (int64, error) {
	res, err := mySQLManager.db.Exec(sql, params...)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (mySQLManager *MySQLManager) Update(sql string, params ...any) (int64, error) {
	res, err := mySQLManager.db.Exec(sql, params...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (mySQLManager *MySQLManager) Delete(sql string, params ...any) (int64, error) {
	res, err := mySQLManager.db.Exec(sql, params...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
