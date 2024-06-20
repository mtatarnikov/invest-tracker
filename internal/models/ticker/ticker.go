package ticker

import (
	"database/sql"
	"invest-tracker/pkg/storage"
)

type Ticker struct {
	Name string
	Type string
}

func List(db storage.Database) ([]string, error) {
	sqlDB := db.Instance()
	if sqlDB == nil {
		return nil, sql.ErrConnDone
	}

	query := `SELECT ticker FROM v_tickers`
	rows, err := sqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickers := []string{""}

	for rows.Next() {
		var ticker Ticker
		if err := rows.Scan(&ticker.Name); err != nil {
			return nil, err
		}
		tickers = append(tickers, ticker.Name)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tickers, nil
}
