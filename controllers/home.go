package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/Yoonnyy/GoMxu/helpers"
	"github.com/go-chi/chi/v5"
)

// GET	/{slug}
func SearchSlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")

	fmt.Println(slug)
}

// POST /
func CreateShortened(w http.ResponseWriter, r *http.Request) {
	// TODO: [Setting] limit file size
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

		// check if url starts with https:// or http://
		if matched, _ := regexp.MatchString("^https?://", url); !matched {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "urls must start with http:// or https:// \n")
			return
		}

		// check blacklisted urls

		// ping url to check if exists

		// generate slug
		// TODO: [Setting] length
		slug := helpers.RandomSlug(6)
		// check if slug already exists

		// save slug to database

		// respond with slug
		io.WriteString(w, fmt.Sprintf("https://localhost:1315/%v\n", slug))
		return
	}

	// shortening url
	if r.MultipartForm.File["file"] != nil {
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		tempFile, err := os.CreateTemp("/tmp", "upload-*.png")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer tempFile.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		tempFile.Write(fileBytes)

		// generate slug
		// TODO: [Setting] length
		slug := helpers.RandomSlug(6)
		// check if slug already exists

		// save slug to database

		// respond with slug
		io.WriteString(w, fmt.Sprintf("https://localhost:1315/%v\n", slug))
		return
	}
}
