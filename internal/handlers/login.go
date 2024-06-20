package handlers

import (
	//"fmt"
	"invest-tracker/internal/models/user"
	"invest-tracker/internal/render"
	"invest-tracker/pkg/storage"
	"net/http"

	"github.com/antonlindstrom/pgstore"
)

func Login(db storage.Database, store *pgstore.PGStore, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		render.RenderTemplate(w, "login.html")
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusInternalServerError)
			return
		}

		user, err := user.GetByLogin(db, r.FormValue("username"))
		if err != nil {
			http.Error(w, "Unauthorized (Wrong username)", http.StatusUnauthorized)
			return
		}

		if !user.CheckPassword(r.FormValue("password")) {
			http.Error(w, "Unauthorized (Invalid password)", http.StatusUnauthorized)
			return
		}

		session, _ := store.Get(r, "session-name")
		session.Values["authenticated"] = true
		session.Values["user_id"] = user.ID
		session.Save(r, w)

		http.Redirect(w, r, "/assets", http.StatusSeeOther)
	}
}
