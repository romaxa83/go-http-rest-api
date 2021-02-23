package store

import (
	"database/sql"
	_ "github.com/lib/pq" // анонимное импортирование
)

type Store struct {
	config *Config
	db *sql.DB
	userRepository *UserRepository
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

// получения UserRepository из нашего store
// если есть - отдаем, если нету , то создаем и отдаем
// из других мест вызываеть store.User().Create()
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}