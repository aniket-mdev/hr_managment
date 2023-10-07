package services

import (
	"database/sql"

	"github.com/aniket-mdev/hr_managment/apis/dto"
	"github.com/aniket-mdev/hr_managment/apis/helper"
	"github.com/aniket-mdev/hr_managment/apis/repositories"
	db "github.com/aniket-mdev/hr_managment/sqlc_lib"
	"github.com/google/uuid"
)

type UserService interface {
	CreateUser(req dto.CreateUserRequestDTO) (db.Users, error)
	GetUserById(user_id string) (db.Users, error)
	ActiveDeactiveUserAccount(req dto.ActiveDeactiveUserAccountRequestDTO) (db.Users, error)
}

type user_service struct {
	user_repo repositories.UserRepository
}

func NewUserService(user_repo repositories.UserRepository) UserService {
	return &user_service{
		user_repo: user_repo,
	}
}

func (user_ser *user_service) CreateUser(req dto.CreateUserRequestDTO) (db.Users, error) {
	args := db.CreateAdminUserParams{
		Name:     req.Name,
		Email:    req.Email,
		Contact:  req.Contact,
		Password: req.Password,
		UserType: req.UserType,
		IsAccountActive: sql.NullBool{
			Bool:  *req.IsAccountActive,
			Valid: true,
		},
	}
	user, err := user_ser.user_repo.CreateUser(args)
	err = helper.HandleDBErr(err)

	return user, err
}

func (user_ser *user_service) GetUserById(user_id string) (db.Users, error) {
	user_obj_id, err := uuid.Parse(user_id)

	if err != nil {
		return db.Users{}, err
	}

	return user_ser.user_repo.GetUserById(user_obj_id)
}

func (user_ser *user_service) ActiveDeactiveUserAccount(req dto.ActiveDeactiveUserAccountRequestDTO) (db.Users, error) {
	user_id, err := uuid.Parse(req.ID)

	if err != nil {
		return db.Users{}, err
	}

	args := db.ActiveDeactiveUserAccountParams{
		ID: user_id,
		IsAccountActive: sql.NullBool{
			Bool:  *req.IsAccountActive,
			Valid: true,
		},
	}

	return user_ser.user_repo.ActiveDeactiveUserAccount(args)
}
