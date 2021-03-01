package sqlstore

import (
	"database/sql"
	"github.com/romaxa83/go-http-rest-api/internal/app/model"
	"github.com/romaxa83/go-http-rest-api/internal/app/store"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) error {

	// валидируем данные
	if err := u.Validate(); err != nil {
		return err
	}

	// запускаем callback перед сохранением
	if err := u.BeforeCreate(); err != nil {
		return err
	}

	// метод Scan будет мапить вернувшийся id в нашу модель юзера
	return r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email, u.EncryptedPassword,
		).Scan(&u.ID)
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = $1",
		email,
		).Scan(
			&u.ID,
			&u.Email,
			&u.EncryptedPassword,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}

			return nil, err
		}

	return u, nil
}