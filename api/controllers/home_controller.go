package controllers

import (
	"net/http"

	"github.com/cigetbudi/fullstack/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Halo jon")
}
