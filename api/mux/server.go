package muxserver

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "mux pong"}`))
	})

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
