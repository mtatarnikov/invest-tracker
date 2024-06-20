package app

import (
	"invest-tracker/internal/handlers"
	"invest-tracker/internal/middleware"
	"invest-tracker/pkg/storage"
	"log"
	"net/http"
	"path/filepath"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/mux"
)

func Run() {
	pg := &storage.PostgresDB{}
	err := pg.Init()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}
	defer pg.Close()

	// Создаем хранилище сессий с использованием pgstore
	store, err := pgstore.NewPGStore("postgres://postgres:p0stgreS@localhost/invest?sslmode=disable", []byte("super-secret-key"))
	if err != nil {
		log.Fatal("Ошибка создания хранилища сессий: ", err)
	}
	defer store.Close()

	r := mux.NewRouter()

	// Настройка обработки статических файлов
	staticPath := filepath.Join("..\\..\\", "ui", "static")
	fs := http.FileServer(http.Dir(staticPath))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// unprotected routes
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		handlers.Login(pg, store, w, r)
	}).Methods("GET", "POST")

	// protected routes
	r.Handle("/", middleware.Auth(http.HandlerFunc(handlers.Home), store)).Methods("GET")
	r.Handle("/page1", middleware.Auth(http.HandlerFunc(handlers.Page1), store)).Methods("GET")
	r.Handle("/assets", middleware.Auth(http.HandlerFunc(handlers.Assets), store)).Methods("GET")
	r.Handle("/assets-table", middleware.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.AssetsTable(pg, store, w, r)
	}), store)).Methods("GET")
	r.Handle("/assets-chart-one", middleware.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.AssetsChartOne(pg, w, r)
	}), store)).Methods("GET")
	r.Handle("/assets-chart-two", middleware.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.AssetsChartTwo(pg, w, r)
	}), store)).Methods("GET")
	r.Handle("/asset-new", middleware.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.AssetNew(pg, store, w, r)
	}), store)).Methods("POST")
	r.Handle("/users", middleware.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.Users(pg, w, r)
	}), store)).Methods("GET")
	r.Handle("/tickers", middleware.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.Tickers(pg, w, r)
	}), store)).Methods("GET")

	http.ListenAndServe(":80", r)
}
