package store

import "github.com/romaxa83/go-http-rest-api/internal/app/model"

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(u *model.User) (*model.User, error) {

	// валидируем данные
	if err := u.Validate(); err != nil {
		return nil, err
	}

	// запускаем callback перед сохранением
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}

	// метод Scan будет мапить вернувшийся id в нашу модель юзера
	if err := r.store.db.QueryRow(
		"INSERT INTO users (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email, u.EncryptedPassword,
		).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
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
			return nil, err
		}

	return u, nil
}