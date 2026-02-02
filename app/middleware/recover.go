package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"vibly/pkg/utils"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("panic recovered: %v", rec)
				var err error
				switch e := rec.(type) {
				case error:
					err = e
				default:
					err = errors.New(fmt.Sprint(e))
				}
				utils.ErrorResponse(w, http.StatusInternalServerError, err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
