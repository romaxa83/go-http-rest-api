package model_test

import (
	"testing"
	"github.com/romaxa83/go-http-rest-api/internal/app/model"

	"github.com/stretchr/testify/assert"
)

// создаем тест в ввиде табличных тестов
// чтоб проверить различные варианты валидации
func TestUser_Validate(t *testing.T) {
	// определяем testCases, в виде массива ананимных структур
	testCases := []struct{
		name string
		u func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "with encrypted password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpassword"
				return u
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = "invalid"
				return u
			},
			isValid: false,
		},
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "short password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "short"
				return u
			},
			isValid: false,
		},
	}

	// итерируем , определеные выше, тест кейсы
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// проверяем , валидный или не валидный кейс, и соответственно проверям должны вернуться ошибка или нет
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
