package db

import "database/sql"

func ConnectDB(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	// TODO: Set Max open conn
	// TODO: Set Max idle conn
	// TODO: Set conn life time

	return db, nil
}
