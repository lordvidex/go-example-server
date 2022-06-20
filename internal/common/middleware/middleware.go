package middleware

import (
	"github.com/lordvidex/go-example-server/internal/common/errors"
	"net/http"
	"strings"
)

// RemoveTrailingSlash is a middleware function that removes trailing slash from a
// URL path before `next http.Handler` is called
func RemoveTrailingSlash(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	}
}

func Recover(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_ = errors.InternalServerError{}.ToJSON(w)
			}
		}()
		next.ServeHTTP(w, r)
	}
}

// Apply TODO: is a convenient function that applies all middleware in the order passed in
func Apply(next ...http.Handler) http.Handler {
	return nil
}
