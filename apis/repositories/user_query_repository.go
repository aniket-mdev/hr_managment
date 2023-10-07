package repositories

import (
	db "github.com/aniket-mdev/hr_managment/sqlc_lib"
	"github.com/google/uuid"
)

func (user_repo *user_repository) CreateUser(args db.CreateAdminUserParams) (db.Users, error) {
	ctx, cancel := user_repo.Init()
	defer cancel()
	return user_repo.db.Queries.CreateAdminUser(ctx, args)
}

func (user_repo *user_repository) GetUserById(user_id uuid.UUID) (db.Users, error) {
	ctx, cancel := user_repo.Init()
	defer cancel()

	return user_repo.db.Queries.GetUserById(ctx, user_id)
}

func (user_repo *user_repository) ActiveDeactiveUserAccount(args db.ActiveDeactiveUserAccountParams) (db.Users, error) {
	ctx, cancel := user_repo.Init()
	defer cancel()
	return user_repo.db.Queries.ActiveDeactiveUserAccount(ctx, args)
}
