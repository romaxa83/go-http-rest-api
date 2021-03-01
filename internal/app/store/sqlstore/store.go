package sqlstore

import (
	"database/sql"
	_ "github.com/lib/pq" // анонимное импортирование
	"github.com/romaxa83/go-http-rest-api/internal/app/store"
)

type Store struct {
	db *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// получения UserRepository из нашего store
// если есть - отдаем, если нету , то создаем и отдаем
// из других мест вызываеть store.User().Create()
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}