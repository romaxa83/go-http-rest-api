package store_test

import (
	"github.com/romaxa83/go-http-rest-api/internal/app/store"
	"github.com/romaxa83/go-http-rest-api/internal/app/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail_Fail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "test@tes.com"
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)
}

func TestUserRepository_FindByEmail_Success(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "test@tes.com"
	_, err := s.User().Create(&model.User{
		Email: email,
		Password: "password",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
