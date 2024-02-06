package server

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/jba/muxpatterns" // Implemented natively in Go 1.22
)

const DEFAULT_PORT = "80"

var tmpl *template.Template = nil

func ListenAndServe() error {
	mux := muxpatterns.NewServeMux()
	mux.HandleFunc("GET /example", getExample)
	tmpl, _ = template.ParseGlob("./templates/*.tmpl.html")

	mux.Handle("/", http.FileServer(http.Dir("public")))

	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}
	slog.Info("FeedGator is listening,", "port", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
}

func getExample(w http.ResponseWriter, r *http.Request) {
	slog.Info("Getting example")

	tmpl.ExecuteTemplate(w, "example", nil)
}
