package handlers

import (
	"net/http"
	"thesis-management-app/views/auth"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, auth.Login())
}
