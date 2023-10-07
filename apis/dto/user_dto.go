package dto

type CreateUserRequestDTO struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Contact         string `json:"contact" binding:"required"`
	Password        string `json:"password" binding:"required"`
	UserType        string `json:"user_type" binding:"required"`
	IsAccountActive *bool  `json:"is_account_active" binding:"required"`
}

type ActiveDeactiveUserAccountRequestDTO struct {
	ID              string `json:"id" binding:"required"`
	IsAccountActive *bool  `json:"is_account_active" binding:"required"`
}
