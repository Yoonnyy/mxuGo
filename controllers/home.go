package controllers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"github.com/Yoonnyy/GoMxu/helpers"
	"github.com/Yoonnyy/GoMxu/models"
	"github.com/go-chi/chi/v5"
)

var (
	slugs models.SlugStore
	urls  models.UrlStore
	files models.FileStore
)

// GET	/{slug}
func SearchSlug(w http.ResponseWriter, r *http.Request) {
	if !config.FileDownloadActive && !config.UrlReditectActive {
		w.WriteHeader(418)
		io.WriteString(w, "services are disabled\n")
		return
	}
	fmt.Println(config.DatabaseUrl)
	// TODO: [Setting] Check if file upload or url shortening is active else reject with message
	slug := chi.URLParam(r, "slug")

	// check if slug exists
	s, err := slugs.GetBySlug(slug)

	if err != nil {
		w.WriteHeader(404)
		io.WriteString(w, "not found\n")
		return
	}

	// handle url redirect
	if !s.IsFile {
		// get the url
		url, err := urls.GetBySlug(slug)

		// if somehow the slug exists but the url doesn't
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, "the url doesn't exist on the server??")
			return
		}

		w.Header().Add("Location", url.Destination)
		w.WriteHeader(301)
		return
	}

	// handle file download
	// TODO: [Setting uploads folder]
	http.ServeFile(w, r, fmt.Sprintf("uploads/%s", s.Slug))
}

// POST /
func CreateShortened(w http.ResponseWriter, r *http.Request) {
	if !config.FileUploadActive && !config.UrlShorteningActive {
		w.WriteHeader(418)
		io.WriteString(w, "services are disabled\n")
		return
	}
	// TODO: [Setting] Check if file upload or url shortening is active else reject with message
	// TODO: [Setting] limit file size
	r.ParseMultipartForm(100 * 1024 * 1024)

	if r.Form.Has("url") && r.MultipartForm.File["file"] != nil {
		// reject if both url and file is present
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "provide a url or a file. not both\n")
		return
	} else if !r.Form.Has("url") && r.MultipartForm.File["file"] == nil {
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
		// TODO: check if slug already exists
		slug := helpers.RandomSlug(6)

		// save slug and the url to database
		err := slugs.Insert(false, slug)
		if err != nil {
			io.WriteString(w, "server error. retry again")
			return
		}

		err = urls.Insert(slug, url, 0)
		if err != nil {
			io.WriteString(w, "server error. retry again")
			return
		}

		// respond with slug
		io.WriteString(w, fmt.Sprintf("https://localhost:1315/%v\n", slug))
		return
	}

	// file upload
	if r.MultipartForm.File["file"] != nil {
		bodyFile, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer bodyFile.Close()

		// read first 512 bytes of the uploaded file
		// to detect ContentType
		first512Byte := make([]byte, 512)
		io.ReadFull(bodyFile, first512Byte)
		fileType := http.DetectContentType(first512Byte)
		_ = fileType
		bodyFile.Seek(0, io.SeekStart)
		// TODO: check forbidden mime types type

		// generate slug
		// TODO: [Setting] length
		slug := helpers.RandomSlug(6)
		// check if slug already exists

		// TODO: [Setting uploads folder]
		// create file
		file, err := os.Create("uploads/" + slug)
		if err != nil {
			io.WriteString(w, "error while uploading file.")
			return
		}
		defer file.Close()

		// read from uploaded file
		fileBytes, err := io.ReadAll(bodyFile)
		if err != nil {
			io.WriteString(w, "error while uploading file.")
			return
		}

		// write bytes to file
		_, err = file.Write(fileBytes)
		if err != nil {
			io.WriteString(w, "error while uploading file.")
			return
		}

		// save slug and the file to database
		slugs.Insert(true, slug)
		files.Insert(handler.Filename, slug, int(handler.Size), 0)

		// respond with slug
		io.WriteString(w, fmt.Sprintf("https://localhost:1315/%v\n", slug))
		return
	}
}
