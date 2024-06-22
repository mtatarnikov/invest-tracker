package user

import (
	"invest-tracker/pkg/storage"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetByLogin(t *testing.T) {
	// Создаем mock DB
	mockDB, err := storage.NewMockDatabase()
	if err != nil {
		t.Fatalf("Ошибка при создании mock базы данных: %v", err)
	}
	defer mockDB.Close()

	// Определяем ожидаемый результат
	expectedUser := User{
		UserSafe: UserSafe{
			ID:    1,
			Name:  "John Doe",
			Login: "johndoe",
		},
		Password: "$2a$10$7EqJtq98hPqEX7fNZaFWoOeL2.gG8IuB6Fh3kGpbVVGFh/G6Qp6/6", // bcrypt hash for "password"
	}

	// Определяем поведение mock
	rows := sqlmock.NewRows([]string{"id", "name", "login", "password"}).
		AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Login, expectedUser.Password)

	mockDB.SqlMock.ExpectQuery(`SELECT id, name, login, password FROM users WHERE login=\$1`).
		WithArgs("johndoe").
		WillReturnRows(rows)

	// Выполняем метод GetByLogin
	user, err := GetByLogin(mockDB, "johndoe")
	if err != nil {
		t.Fatalf("Ошибка при вызове GetByLogin: %v", err)
	}

	// Проверяем результат
	assert.Equal(t, expectedUser.ID, user.ID, "ID не совпадает")
	assert.Equal(t, expectedUser.Name, user.Name, "Name не совпадает")
	assert.Equal(t, expectedUser.Login, user.Login, "Login не совпадает")
	assert.Equal(t, expectedUser.Password, user.Password, "Password не совпадает")

	// Проверяем ожидания mock
	if err := mockDB.SqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("Не все ожидания были выполнены: %v", err)
	}
}
