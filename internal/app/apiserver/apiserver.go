package apiserver

import (
	"database/sql"
	"github.com/romaxa83/go-http-rest-api/internal/app/store/sqlstore"
	"net/http"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}