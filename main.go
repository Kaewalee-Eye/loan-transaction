package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Resp struct {
	Ok  bool   `json:"ok"`
	Now string `json:"now,omitempty"`
	Err string `json:"err,omitempty"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	log.Println("DB connected")

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, Resp{Ok: true, Now: time.Now().Format(time.RFC3339)})
	})

	mux.HandleFunc("/inquiry", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		var now time.Time
		if err := db.QueryRowContext(ctx, "select now()").Scan(&now); err != nil {
			writeJSON(w, Resp{Ok: false, Err: err.Error()})
			return
		}
		writeJSON(w, Resp{Ok: true, Now: now.Format(time.RFC3339)})
	})

	log.Println("listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}
