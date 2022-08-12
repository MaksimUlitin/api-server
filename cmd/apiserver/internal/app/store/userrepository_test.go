package store_test

import (
	"testing"

	"github.com/MaksimUlitin/api-server/cmd/apiserver/internal/app/model"
	"github.com/MaksimUlitin/api-server/cmd/apiserver/internal/app/store"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Creat(t *testing.T) {
	s, teardewn := store.TestStore(t, DataBaseURL)
	defer teardewn("users")
	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository(t *testing.T) {
	s, teardewn := store.TestStore(t, DataBaseURL)
	defer teardewn("users")
	Email := "user@exemple.org"
	_, err := s.user.FindByEmail(Email)
	assert.Error(t, err)

	s.user().Create(&model.User{
		Email: "user@example.org",
	})

	u, err := s.user.FindByEmail(Email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
