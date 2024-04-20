package standardserver

import "net/http"

func Start(port string) {
	r := http.NewServeMux()

	r.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "standard pong"}`))
	})

	http.ListenAndServe(port, r)
}
