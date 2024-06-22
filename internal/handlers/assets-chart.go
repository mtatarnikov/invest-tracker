package handlers

import (
	"encoding/json"
	chart "invest-tracker/internal/models/asset"
	"invest-tracker/pkg/storage"
	"log"
	"net/http"

	"github.com/antonlindstrom/pgstore"
)

// JSON для заполнения диаграммы "Аналитика. Активы", которая отображается в assets-list.html
func AssetsChartOne(db storage.Database, store *pgstore.PGStore, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, _ := store.Get(r, "session-name")

		userID, ok := session.Values["user_id"].(int)
		if !ok {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}
		data, err := chart.GetChartDataOne(db.Instance(), userID)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

// JSON для заполнения диаграммы "Аналитика. Активы", которая отображается в assets-list.html
func AssetsChartTwo(db storage.Database, store *pgstore.PGStore, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		session, _ := store.Get(r, "session-name")

		userID, ok := session.Values["user_id"].(int)
		if !ok {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}
		data, err := chart.GetChartDataTwo(db.Instance(), userID)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
