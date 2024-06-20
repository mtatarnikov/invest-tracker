package user

import (
	"database/sql"
	"fmt"
	"invest-tracker/pkg/storage"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserSafe
	Password string
}

type UserSafe struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Login string `json:"login"`
}

type UserWithLinks struct {
	UserSafe
	LinkEdit   string `json:"link_edit"`
	LinkDelete string `json:"link_delete"`
}

func GetByLogin(db storage.Database, login string) (*User, error) {
	sqlDB := db.Instance()
	if sqlDB == nil {
		return nil, sql.ErrConnDone
	}

	var user User

	query := `SELECT id, name, login, password FROM users WHERE login=$1`
	err := sqlDB.QueryRow(query, login).Scan(&user.ID, &user.Name, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func List(db storage.Database) ([]UserWithLinks, error) {
	sqlDB := db.Instance()
	if sqlDB == nil {
		return nil, sql.ErrConnDone
	}

	var usersWithLinks []UserWithLinks

	query := `SELECT id, name, login FROM users`
	rows, err := sqlDB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user UserSafe
		if err := rows.Scan(&user.ID, &user.Name, &user.Login); err != nil {
			return nil, err
		}
		userWithLinks := UserWithLinks{
			UserSafe:   user,
			LinkEdit:   fmt.Sprintf("/user/%d/edit", user.ID),
			LinkDelete: fmt.Sprintf("/user/%d/delete", user.ID),
		}
		usersWithLinks = append(usersWithLinks, userWithLinks)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usersWithLinks, nil
}
