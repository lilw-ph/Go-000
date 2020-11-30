package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type queryempty interface {
	Empty() bool // query result is empty
}

type SqlError struct {
	Err     string // error msg
	IsEmpty bool   // query result is empty
}

func (e *SqlError) Error() string {
	return e.Err
}

func (e *SqlError) Empty() bool {
	return e.IsEmpty
}

func IsQueryEmpty(err error) bool {
	qe, ok := err.(queryempty)
	return ok && qe.Empty()
}

func QueryRow(db *sql.DB, query string, args ...interface{}) (string, error) {
	var (
		e      error = nil
		result string
	)
	err := db.QueryRow(query, args).Scan(&result)
	if err == sql.ErrNoRows {
		e = errors.Wrap(&SqlError{
			Err:     err.Error(),
			//Err:     "query error", // test
			IsEmpty: true,
		}, query)
	}

	return result, e
}
