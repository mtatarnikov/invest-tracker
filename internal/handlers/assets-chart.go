package handlers

import (
	"encoding/json"
	chart "invest-tracker/internal/models/asset"
	"invest-tracker/pkg/storage"
	"net/http"
)

// JSON для заполнения диаграммы "Аналитика. Активы", которая отображается в assets-list.html
func AssetsChartOne(db storage.Database, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := chart.GetChartDataOne()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

// JSON для заполнения диаграммы "Аналитика. Активы", которая отображается в assets-list.html
func AssetsChartTwo(db storage.Database, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data := chart.GetChartDataTwo()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}
