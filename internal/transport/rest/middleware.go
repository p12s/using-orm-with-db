package rest

import (
	"net/http"
)

// loggingMiddleware - logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logging-example, if it need:
		// _, _ = fmt.Fprintf(os.Stdout, "%s: [%s] - %s ", time.Now().Format(time.RFC3339), r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
