package teststore_test

import (
	"github.com/romaxa83/go-http-rest-api/internal/app/model"
	"github.com/romaxa83/go-http-rest-api/internal/app/store"
	"github.com/romaxa83/go-http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {

	s := teststore.New()
	u := model.TestUser(t)
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail_Fail(t *testing.T) {

	s := teststore.New()
	email := "test@tes.com"
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
}

func TestUserRepository_FindByEmail_Success(t *testing.T) {

	s := teststore.New()
	email := "test@tes.com"
	s.User().Create(&model.User{
		Email: email,
		Password: "password",
	})
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
