package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
}

func main() {
	app := application{}

	server := &http.Server{
		Addr:    ":1315",
		Handler: app.routes(),
	}

	server.ListenAndServe()
}

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{slug}", searchSlug)
	r.
		With(middleware.AllowContentType("multipart/form-data")).
		Post("/", createShortened)

	return r
}

// GET	/{slug}
func searchSlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	fmt.Println(slug)
}

// POST /
func createShortened(w http.ResponseWriter, r *http.Request) {
	// TODO: limit file size
	r.ParseMultipartForm(100 * 1024 * 1024)

	if r.Form.Has("url") && r.MultipartForm.File["file"] != nil {
		// reject if both url and file is present
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "provide a url or a file. not both\n")
		return
	} else if !r.Form.Has("url") || r.MultipartForm.File["file"] != nil {
		// reject if both url and file isn't present
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "provide a url or a file\n")
		return
	}

	// shortening url
	if r.Form.Has("url") {
		// get the url
		url := r.FormValue("url")

		if matched, _ := regexp.MatchString(url, "^https?://"); !matched {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "urls must start with http:// or https:// \n")
			return
		}

		// generate slug
	}
	// file := r.FormValue("file")

	// file, header, err := r.FormFile("file")

	// if err != nil {

	// }
	// respond with slug
}

func randomSlug(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}
