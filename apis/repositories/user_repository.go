package repositories

import (
	"context"
	"time"

	"github.com/aniket-mdev/hr_managment/apis"
	db "github.com/aniket-mdev/hr_managment/sqlc_lib"
	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(db.CreateAdminUserParams) (db.Users, error)
	GetUserById(uuid.UUID) (db.Users, error)
	ActiveDeactiveUserAccount(db.ActiveDeactiveUserAccountParams) (db.Users, error)
}

type user_repository struct {
	db *apis.Store
}

func NewUserRepository(db *apis.Store) UserRepository {
	return &user_repository{
		db: db,
	}
}

func (user_repo *user_repository) Init() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}
