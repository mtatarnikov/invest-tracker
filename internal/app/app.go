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

func initDB() (*storage.PostgresDB, error) {
	pg := &storage.PostgresDB{}
	err := pg.Init()
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}
	return pg, nil
}

func readConfig() (config.Config, error) {
	cfg, err := config.Read()
	if err != nil {
		return config.Config{}, fmt.Errorf("failed to read config: %w", err)
	}
	return cfg, nil
}

func initSessionStore(cfg config.Config) (*pgstore.PGStore, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.DBName)

	store, err := pgstore.NewPGStore(dbURL, []byte("super-secret-key"))
	if err != nil {
		return nil, fmt.Errorf("ошибка создания хранилища сессий: %w", err)
	}
	return store, nil
}

func setupRouter(pg *storage.PostgresDB, store *pgstore.PGStore) *mux.Router {
	r := mux.NewRouter()

	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}
	staticPath := filepath.Join(cfg.HtmlUiStaticPath, "ui", "static")
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
		handlers.AssetsChartOne(pg, store, w, r)
	}), store)).Methods("GET")
	r.Handle("/assets-chart-two", middleware.Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.AssetsChartTwo(pg, store, w, r)
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

	return r
}

func Run() {
	pg, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer pg.Close()

	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	store, err := initSessionStore(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	r := setupRouter(pg, store)

	log.Fatal(http.ListenAndServe(":80", r))
}
