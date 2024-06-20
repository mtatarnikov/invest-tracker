package handlers

import (
	"fmt"
	"invest-tracker/internal/models/asset"
	"invest-tracker/pkg/config"
	"invest-tracker/pkg/storage"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/antonlindstrom/pgstore"
)

type MockAssetStore struct {
	WasSaveCalled bool
}

func (m *MockAssetStore) Save(db storage.Database, a asset.Asset) error {
	m.WasSaveCalled = true
	return nil
}

func TestAssetNew(t *testing.T) {
	// Создаем экземпляр структуры, которая реализует интерфейс Database
	db := &storage.MockDatabase{}
	mockStore := &MockAssetStore{}

	// Устанавливаем мок хранилище
	SetAssetStore(mockStore)

	// Создаем форму с данными для отправки
	formData := url.Values{}
	formData.Set("ticker", "AAAA")
	formData.Set("transaction_date", "2020-01-02")
	formData.Set("transaction_type", "buy")
	formData.Set("amount", "10")
	formData.Set("price", "150.50")
	formData.Set("tax", "15.05")
	formData.Set("note", "Long term investment")

	// Создаем запрос с формой
	req, err := http.NewRequest("POST", "/asset-new", strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Чтение конфигурации
	cfg, err := config.Read()
	if err != nil {
		t.Fatal(err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	// Создаем хранилище сессий
	store, err := pgstore.NewPGStore(dsn, []byte("super-secret-key"))
	if err != nil {
		t.Fatal(err)
	}
	defer store.Close()

	// Создаем ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Создаем сессию и устанавливаем user_id
	session, _ := store.Get(req, "session-name")
	session.Values["user_id"] = 1 // Устанавливаем user_id в значение 1 (или любое другое валидное значение)
	session.Save(req, rr)

	// Вызываем метод AssetNew
	AssetNew(db, store, rr, req)

	// Проверяем статус код ответа
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}

	// Проверяем, что в базу данных был добавлен новый актив
	if !mockStore.WasSaveCalled {
		t.Errorf("expected Asset.Save to be called")
	}
}
