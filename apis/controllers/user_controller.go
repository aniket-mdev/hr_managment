package controllers

import (
	"net/http"

	"github.com/aniket-mdev/hr_managment/apis/dto"
	"github.com/aniket-mdev/hr_managment/apis/services"
	"github.com/aniket-mdev/hr_managment/logger"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(*gin.Context)
}

type user_controller struct {
	user_ser services.UserService
}

func NewUserController(user_ser services.UserService) UserController {
	return &user_controller{
		user_ser: user_ser,
	}
}

func (user_con *user_controller) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequestDTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":   "failed to process",
			"error": err,
		})
		return
	}

	logger.DebugLogger.Println("\n", req)

	user, err_ := user_con.user_ser.CreateUser(req)

	if err_ != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":   "failed to process",
			"error": err_.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":    "user account has been created",
		"user":   user,
		"status": true,
	})

}
