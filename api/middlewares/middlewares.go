package middlewares

import (
	"errors"
	"net/http"

	"blog-application/api/auth"
	"blog-application/api/responses"
)

// ensure that all HTTP responses sent by the application have a Content-Type header set to application/json. This middleware can be used to enforce this header for a set of handlers in the application, helping to ensure that responses are properly formatted and easy to process.
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// used as a middleware to check the authentication of a request
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
