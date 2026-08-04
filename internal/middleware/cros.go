package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Allow your frontend
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")

		// Allowed methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Allowed headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Allow cookies if needed
		// w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
