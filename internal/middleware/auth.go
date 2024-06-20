package middleware

import (
	"net/http"

	"github.com/antonlindstrom/pgstore"
)

func Auth(next http.Handler, store *pgstore.PGStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")

		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			// http.Error(w, "Forbidden", http.StatusForbidden)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	})
}
