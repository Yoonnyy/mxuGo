package main

import (
	"net/http"
	"os"

	"github.com/Yoonnyy/GoMxu/db"
	"github.com/Yoonnyy/GoMxu/models"
	_ "github.com/lib/pq"
)

// TODO: convert stores to interfaces
type application struct {
	Models models.Models
}

func (app *application) Serve() error {
	server := &http.Server{
		Addr:    ":1315",
		Handler: app.routes(),
	}

	return server.ListenAndServe()
}

func main() {
	// database connection
	db, err := db.ConnectDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var app = application{
		Models: models.Init(db),
	}

	app.Serve()
}
