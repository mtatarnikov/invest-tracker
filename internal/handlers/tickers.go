package handlers

import (
	"encoding/json"
	ticker "invest-tracker/internal/models/ticker"
	"invest-tracker/pkg/storage"
	"net/http"
)

func Tickers(db storage.Database, w http.ResponseWriter, r *http.Request) {
	tickers, err := ticker.List(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tickers)
}
