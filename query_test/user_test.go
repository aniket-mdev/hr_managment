package querytest

import (
	"context"
	"database/sql"
	"testing"

	db "github.com/aniket-mdev/hr_managment/sqlc_lib"
	"github.com/aniket-mdev/hr_managment/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func create_random_user(t *testing.T) db.Users {
	args := db.CreateAdminUserParams{
		Name:     utils.RandomUserName(6),
		Email:    utils.RandomUserEmail(),
		Contact:  utils.RandomUserContact(),
		Password: "test123",
		UserType: "Admin",
		IsAccountActive: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	}

	user, err := testQueries.CreateAdminUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Name, args.Name)
	require.Equal(t, user.Email, args.Email)
	require.Equal(t, user.Password, args.Password)
	require.Equal(t, user.Contact, args.Contact)
	require.Equal(t, user.UserType, args.UserType)
	require.Equal(t, user.IsAccountActive, args.IsAccountActive)

	return user
}

func TestCreateUser(t *testing.T) {
	create_random_user(t)
}

func TestGetUserById(t *testing.T) {
	user_id := "e4839d69-0177-4ba6-8149-8d31ab488ee3"

	user_obj_id, err := uuid.Parse(user_id)

	require.NoError(t, err)

	user, err := testQueries.GetUserById(context.Background(), user_obj_id)

	require.NoError(t, err)
	require.Equal(t, user.ID, user_obj_id)
	require.NotEmpty(t, user)
}

func TestActiveDeactiveUserAccount(t *testing.T) {
	user_id := "e4839d69-0177-4ba6-8149-8d31ab488ee3"

	user_obj_id, err := uuid.Parse(user_id)

	require.NoError(t, err)
	args := db.ActiveDeactiveUserAccountParams{
		ID: user_obj_id,
		IsAccountActive: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	}

	user, err := testQueries.ActiveDeactiveUserAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user_obj_id, user.ID)
}
