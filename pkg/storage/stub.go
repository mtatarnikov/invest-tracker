package storage

// import (
// 	"database/sql"
// )

// // Объявляем интерфейс DataRetriever, который определяет методы для получения данных
// type DataRetriever interface {
// 	QueryRow(query string, args ...interface{}) *sql.Row
// }

// // StubDatabase содержит мапу с анонимными функциями для различных запросов
// type StubDatabase struct {
// 	Queries map[string]func(args ...interface{}) *sql.Row
// }

// // NewStubDatabase создает новый экземпляр StubDatabase
// func NewStubDatabase() *StubDatabase {
// 	return &StubDatabase{
// 		Queries: make(map[string]func(args ...interface{}) *sql.Row),
// 	}
// }

// // QueryRow вызывает функцию, соответствующую запросу, если она есть в мапе
// func (s *StubDatabase) QueryRow(query string, args ...interface{}) *sql.Row {
// 	if queryFunc, exists := s.Queries[query]; exists {
// 		return queryFunc(args...)
// 	}
// 	return nil
// }

// func (s *StubDatabase) Init() error {
// 	return nil
// }

// func (s *StubDatabase) Instance() *sql.DB {
// 	return nil
// }

// func (s *StubDatabase) Close() error {
// 	return nil
// }
