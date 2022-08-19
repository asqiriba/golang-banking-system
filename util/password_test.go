package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(5)
	hashedPassword, err := HashedPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = CheckPassword(hashedPassword, password)
	require.NoError(t, err)

	wrongPassword := RandomString(5)
	err = CheckPassword(hashedPassword, wrongPassword)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
