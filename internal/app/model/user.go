package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)
/*
 ключ "omitempty", говорит что если поле пустое, то не возвращать в ответе
 ключ "-", говорит что вообще не возвращать при рендере
*/
type User struct {
	ID 				  int `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"password,omitempty"`
	EncryptedPassword string `json:"-"`
}

func (u *User) Validate() error {
	// для password создаем кастомный валидатор, так как при редактировании
	// наша модели будет не валидна из-за того что пароль храним в другом поле
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(requiredIf(u.EncryptedPassword == "")), validation.Length(6, 100)),
	)
}

// метод вызываеться каждый раз когда мы сохраняем пользователя
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := encryptString(u.Password)
		if err != nil {
			return nil
		}

		u.EncryptedPassword = enc
	}

	return nil
}

// функция затирает поля, которые мы не хотим отдавать
func (u *User) Sanitize() {
	u.Password = ""
}
// сравнение зашифрованого пароля в бд и введеного пользователем
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

func encryptString(s string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}