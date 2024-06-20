package asset

import (
	"database/sql"
	"fmt"
	"invest-tracker/pkg/storage"
	"log"
	"time"
)

type Asset struct {
	Id              int            `json:"id"`
	UserId          int            `json:"user_id"`
	Ticker          string         `json:"ticker"`
	TransactionType string         `json:"transaction_type"`
	TransactionDate string         `json:"transaction_date"`
	Amount          int            `json:"amount"`
	Price           float64        `json:"price"`
	Tax             float64        `json:"tax"`
	Note            sql.NullString `json:"note"`
}

type AssetsWithLinks struct {
	Asset
	LinkEdit   string `json:"link_edit"`
	LinkDelete string `json:"link_delete"`
}

func List(db storage.Database, userID int) ([]AssetsWithLinks, error) {
	sqlDB := db.Instance()
	if sqlDB == nil {
		return nil, sql.ErrConnDone
	}

	var assetsWithLinks []AssetsWithLinks

	query := `SELECT id, ticker, transaction_type, transaction_date, amount, price, tax, note
	            FROM assets
			   WHERE user_id = $1`
	rows, err := sqlDB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var asset Asset
		var note sql.NullString
		var transactionDate time.Time

		if err := rows.Scan(&asset.Id, &asset.Ticker, &asset.TransactionType, &transactionDate, &asset.Amount, &asset.Price, &asset.Tax, &note); err != nil {
			return nil, err
		}

		asset.TransactionDate = transactionDate.Format("2006-01-02") // Форматирование даты (операции с активом)
		asset.Note = note
		assetWithLinks := AssetsWithLinks{
			Asset:      asset,
			LinkEdit:   fmt.Sprintf("/asset/%d/edit", asset.Id),
			LinkDelete: fmt.Sprintf("/asset/%d/delete", asset.Id),
		}
		assetsWithLinks = append(assetsWithLinks, assetWithLinks)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return assetsWithLinks, nil
}

// Save сохраняет актив в базу данных.
func Save(db storage.Database, a Asset) error {
	sqlDB := db.Instance()
	if sqlDB == nil {
		log.Println("Подключение к базе данных не установлено")
		return sql.ErrConnDone
	}

	tx, err := sqlDB.Begin()
	if err != nil {
		log.Println("Не удалось начать транзакцию:", err)
		return err
	}

	query := `INSERT INTO assets (user_id, ticker, transaction_type, transaction_date, amount, price, tax, note)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	log.Printf("Выполняем запрос: %s\n", query)
	if _, err := tx.Exec(query, a.UserId, a.Ticker, a.TransactionType, a.TransactionDate, a.Amount, a.Price, a.Tax, a.Note); err != nil {
		tx.Rollback()
		log.Printf("Ошибка при выполнении запроса: %s\n", err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Println("Не удалось зафиксировать транзакцию:", err)
		return err
	}

	return nil
}
