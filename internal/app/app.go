package app

import (
	"fmt"
	"invest-tracker/internal/handlers"
	"invest-tracker/internal/middleware"
	"invest-tracker/pkg/config"
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
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("failed to read config: %w", err)
	}

	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.DBName)

	store, err := pgstore.NewPGStore(dbURL, []byte("super-secret-key"))
	if err != nil {
		log.Fatal("Ошибка создания хранилища сессий: ", err)
	}
	defer store.Close()

	r := mux.NewRouter()

	// Настройка обработки статических файлов
	// Для локальной разработки
	//staticPath := filepath.Join("..\\..\\", "ui", "static")
	staticPath := filepath.Join("ui", "static")
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
