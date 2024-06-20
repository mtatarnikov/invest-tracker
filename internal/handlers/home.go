package handlers

import (
	"fmt"
	//"myapp/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is HOME")
	//render.RenderTemplate(w, "home.page.html")
}
