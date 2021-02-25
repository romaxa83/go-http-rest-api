package model

import "testing"

// хелпер, который возвращает валидные данные для юзера

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Email: "test@test.com",
		Password: "password",
	}
}