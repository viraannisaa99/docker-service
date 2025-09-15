package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type resp struct {
	Message string `json:"message"` // memberi nama field JSON explicitly saat di encode
	Env     string `json:"env"` 
}

func main() {
	mux := http.NewServeMux() // router minimalis untuk server kecil

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"}) // abaikan error dari encode
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp{
			Message: "Hello Vira!",
			Env:     os.Getenv("GO_ENV"), // ambil dari environment variable GO_ENV
		})
	})

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		type Ping struct {
			Pong bool   `json:"pong"`
			Time string `json:"time"`
		}

		if err := json.NewEncoder(w).Encode(Ping{
			Pong: true,
			Time: time.Now().Format(time.RFC3339),
		}); err != nil {
			// optional: log error biar ketauan kalau encoding gagal
			log.Printf("encode error: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
	})

	addr := ":8080"
	log.Printf("listening on %s ...", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
