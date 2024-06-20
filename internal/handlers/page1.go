package handlers

import (
	"fmt"
	"net/http"
)

func Page1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Page #1")
}
