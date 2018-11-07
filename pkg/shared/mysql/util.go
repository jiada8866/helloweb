package mysql

import (
	"database/sql"

	"github.com/pkg/errors"
)

func Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := DB.Exec(query, args...)
	if err != nil {
		return result, errors.Wrapf(err, "failed to exec with sql=%s, args=%v", query, args)
	}
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := DB.Query(query, args...)
	if err != nil {
		return rows, errors.Wrapf(err, "failed to query with sql=%s, args=%v", query, args)
	}
	return rows, err
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	return DB.QueryRow(query, args...)
}

func Prepare(query string) (*sql.Stmt, error) {
	stmt, err := DB.Prepare(query)
	if err != nil {
		return stmt, errors.Wrapf(err, "failed to prepare with sql=%s", query)
	}
	return stmt, err
}
