package handlers

import (
	"net/http"
	"thesis-management-app/views/home"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
