package controllers

import (
	"net/http"

	"github.com/aniket-mdev/hr_managment/apis/dto"
	"github.com/aniket-mdev/hr_managment/apis/helper"
	"github.com/aniket-mdev/hr_managment/apis/services"
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
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, err.Error(), helper.EmptyObj{}, helper.USER_DATA)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	user, err := user_con.user_ser.CreateUser(req)

	if err != nil {
		response := helper.BuildFailedResponse(helper.FAILED_PROCESS, err.Error(), helper.EmptyObj{}, helper.USER_DATA)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.BuildSuccessResponse(helper.USER_REGISTRATION_SUCCESS, user, helper.USER_DATA)
	ctx.JSON(http.StatusOK, response)

}
