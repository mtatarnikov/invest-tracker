package storage

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockDatabase struct {
	SqlDB        *sql.DB
	SqlMock      sqlmock.Sqlmock
	QueryRowFunc func(query string, args ...interface{}) *sql.Row
}

func NewMockDatabase() (*MockDatabase, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}
	return &MockDatabase{
		SqlDB:   db,
		SqlMock: mock,
	}, nil
}

func (db *MockDatabase) Init() error {
	return nil
}

func (db *MockDatabase) Instance() *sql.DB {
	return db.SqlDB
}

func (db *MockDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
	if db.QueryRowFunc != nil {
		return db.QueryRowFunc(query, args...)
	}
	return db.SqlDB.QueryRow(query, args...)
}

func (db *MockDatabase) Close() error {
	return db.SqlDB.Close()
}
