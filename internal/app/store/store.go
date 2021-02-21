package store

import (
	"database/sql"
	_ "github.com/lib/pq" // анонимное импортирование
)

type Store struct {
	config *Config
	db *sql.DB
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}
// подключаемся к бд
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURL)

	if err != nil {
		return nil
	}
	// пингуем подключение
	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

// закрываем соединение с бд
func (s *Store) Close() {
	s.db.Close()
}