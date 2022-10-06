package middlewares

import (
	"errors"
	"net/http"

	"github.com/cigetbudi/fullstack/api/auth"
	"github.com/cigetbudi/fullstack/api/responses"
)

func SetMiddlewareJson(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content_Type", "application/json")
		next(w, r)
	}
}
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		}
		next(w, r)
	}

}
