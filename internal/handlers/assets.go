package handlers

import (
	"database/sql"
	"encoding/json"
	"invest-tracker/internal/models/asset"
	"invest-tracker/internal/render"
	"invest-tracker/pkg/storage"
	"net/http"
	"strconv"
	"time"

	"github.com/antonlindstrom/pgstore"
)

var assetStore asset.AssetStore = &asset.RealAssetStore{}

func SetAssetStore(store asset.AssetStore) {
	assetStore = store
}

func Assets(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		render.RenderTemplate(w, "assets-list.html")
		return
	}
}

// JSON для заполнения таблицы "Активы", которая отображается в assets-list.html
func AssetsTable(db storage.Database, store *pgstore.PGStore, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		session, _ := store.Get(r, "session-name")

		userID, ok := session.Values["user_id"].(int)
		if !ok {
			http.Error(w, "User not authenticated", http.StatusUnauthorized)
			return
		}

		assets, err := asset.List(db, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(assets)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}

func AssetNew(db storage.Database, store *pgstore.PGStore, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Метод не поддерживается.", http.StatusMethodNotAllowed)
		return
	}

	session, _ := store.Get(r, "session-name")

	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Ошибка при разборе формы.", http.StatusInternalServerError)
		return
	}

	// Извлечение данных из формы.
	ticker := r.FormValue("ticker")
	transactionDate := r.FormValue("transaction_date")
	_, err := time.Parse("2006-01-02", transactionDate)
	if err != nil {
		http.Error(w, "Неверный формат даты", http.StatusBadRequest)
		return
	}
	transactionType := r.FormValue("transaction_type")
	amount, err := strconv.Atoi(r.FormValue("amount"))
	if err != nil {
		http.Error(w, "Неверный формат количества.", http.StatusBadRequest)
		return
	}
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Error(w, "Неверный формат цены.", http.StatusBadRequest)
		return
	}
	taxStr := r.FormValue("tax")
	var tax float64
	if taxStr != "" {
		tax, err = strconv.ParseFloat(taxStr, 64)
		if err != nil {
			http.Error(w, "Неверный формат налога.", http.StatusBadRequest)
			return
		}
	}
	note := r.FormValue("note")

	newAsset := asset.Asset{
		UserId:          userID,
		Ticker:          ticker,
		TransactionType: transactionType,
		TransactionDate: transactionDate,
		Amount:          amount,
		Price:           price,
		Tax:             tax,
		Note:            sql.NullString{String: note, Valid: note != ""},
	}

	// Сохранение нового актива в базу данных.
	if err := assetStore.Save(db, newAsset); err != nil {
		http.Error(w, "Ошибка при сохранении актива.", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/assets", http.StatusSeeOther)
}
