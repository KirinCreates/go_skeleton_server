package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"goServerSkeleton/templates"
)

func getTimeHandler(w http.ResponseWriter, r *http.Request) {
	response := templates.SimpleResponse("Server time: " + time.Now().String())
	response.Render(r.Context(), w)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	exampleContent := templates.ExampleContent("Click me!")
	base := templates.Base("Learn to Pro", "Go fullstack skeleton project: Go+Templ+HTMX (GoTH-stack)", exampleContent)
	base.Render(r.Context(), w)
}

func fileServer(r chi.Router, path string) {
	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err.Error())
	}

	baseAndTargetPath := filepath.Join(root, path)

	path += "/*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())

		GETPathRegistered := rctx.RoutePattern()
		pathPrefix := strings.TrimSuffix(GETPathRegistered, "/*")                        //removes trailing /* if it exists
		fs := http.StripPrefix(pathPrefix, http.FileServer(http.Dir(baseAndTargetPath))) //remove pathPrefix (/static) from request and append the rest to fs root serving folder

		fs.ServeHTTP(w, r)
	})
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.ClientIPFromHeader("CF-Connecting-IP")) //if you use cloudflare. Remove otherwise
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(5 * time.Second))

	server := &http.Server{
		Addr:         "127.0.0.1:8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	fileServer(r, "/static")

	//all routes for website
	r.Get("/", homeHandler)
	r.Get("/getTime", getTimeHandler)

	server.ListenAndServe()
}
