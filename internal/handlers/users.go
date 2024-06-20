package handlers

import (
	"encoding/json"
	"invest-tracker/internal/models/user"
	"invest-tracker/pkg/storage"
	"net/http"
)

func Users(db storage.Database, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		users, err := user.List(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		jsonResponse, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResponse)
	}
}
