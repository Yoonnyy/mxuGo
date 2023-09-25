package main

import (
	"fmt"
	"net/http"

	"github.com/Yoonnyy/GoMxu/configuration"
	"github.com/Yoonnyy/GoMxu/db"
	"github.com/Yoonnyy/GoMxu/models"
	_ "github.com/lib/pq"
)

// TODO: convert stores to interfaces
type application struct {
	Models models.Models
	Config configuration.Config
}

func (app *application) Serve() error {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", app.Config.Address, app.Config.Port),
		Handler: app.routes(),
	}

	return server.ListenAndServe()
}

func main() {
	var app = application{
		Config: *configuration.ParseConfig(),
	}
	// database connection
	db, err := db.ConnectDB(app.Config.DatabaseUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	app.Models = models.Init(db)

	app.Serve()
}
