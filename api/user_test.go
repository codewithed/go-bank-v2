package api

import (
	"testing"

	db "github.com/codewithed/go-bank-v2/db/sqlc"
	"github.com/codewithed/go-bank-v2/util"
	"github.com/stretchr/testify/require"
)

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.CreatePasswordHash(password)
	require.NoError(t, err)

	user = db.User{
		Username:     util.RandomOwner(),
		PasswordHash: hashedPassword,
		FullName:     util.RandomOwner(),
		Email:        util.RandomEmail(),
	}
	return
}
