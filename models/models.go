package models

import "database/sql"

type Models struct {
	Slugs SlugStore
	Urls  UrlStore
	Files FileStore
}

var db *sql.DB

func Init(conn *sql.DB) Models {
	db = conn
	return Models{}
}
