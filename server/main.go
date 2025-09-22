package main

import (
	"fmt"
	"html"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"path": %q}`, html.EscapeString(r.URL.Path))
	})

	slog.Info("Starting server on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
