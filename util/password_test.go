package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(8)

	passwordHash1, err := CreatePasswordHash(password)
	require.NoError(t, err)
	require.NotEmpty(t, passwordHash1)

	err = ValidatePassword(password, passwordHash1)
	require.NoError(t, err)

	wrongPassword := RandomString(7)
	err = ValidatePassword(wrongPassword, passwordHash1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	passwordHash2, err := CreatePasswordHash(password)
	require.NoError(t, err)
	require.NotEmpty(t, passwordHash2)
	require.NotEqual(t, passwordHash1, passwordHash2)
}
