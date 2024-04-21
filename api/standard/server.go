package standardserver

import (
	"log"
	"net/http"
)

func Start(port string) {
	r := http.NewServeMux()

	r.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "standard pong"}`))
	})

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
