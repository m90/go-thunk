package thunk

import (
	"net/http"
)

// HandleSafelyWith wraps the passed handler in a defer / recover block
// preventing the handler from crashing the application on panics
// and passing eventual errors to the given callback
func HandleSafelyWith(callback func(error)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			RunSafelyWith(func() {
				next.ServeHTTP(w, r)
			}, func(err error) {
				callback(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			})
		})
	}
}

// HandleSafely wraps the passed handler in a defer / recover block
// preventing the handler from crashing the application on panics
// while discarding potential errors
func HandleSafely() func(http.Handler) http.Handler {
	return HandleSafelyWith(func(_ error) {})
}
