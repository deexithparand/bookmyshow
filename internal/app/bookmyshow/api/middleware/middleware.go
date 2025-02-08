package middleware

import (
	"fmt"
	"net/http"

	"github.com/deexithparand/bookmyshow/internal/app/bookmyshow/customerrors"
)

func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				customerrors.ServerErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
