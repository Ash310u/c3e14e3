package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", pingHandler)
	mux.HandleFunc("/schema", authRequiredHandler)
	mux.HandleFunc("/extract", authRequiredHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	log.Printf("server listening on %s", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "pong")
}

func authRequiredHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/schema":
		if r.Method != http.MethodPost {
			methodNotAllowed(w)
			return
		}
		writeJSONError(w, http.StatusUnauthorized, "authentication required")
		return
	case "/extract":
		if r.Method != http.MethodPost {
			methodNotAllowed(w)
			return
		}
		writeJSONError(w, http.StatusUnauthorized, "authentication required")
		return
	default:
		writeJSONError(w, http.StatusNotFound, "not found")
		return
	}
}

func methodNotAllowed(w http.ResponseWriter) {
	writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
}

func writeJSONError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}
