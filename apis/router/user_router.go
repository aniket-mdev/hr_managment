package router

import (
	"github.com/aniket-mdev/hr_managment/apis"
	"github.com/aniket-mdev/hr_managment/apis/controllers"
	"github.com/aniket-mdev/hr_managment/apis/repositories"
	"github.com/aniket-mdev/hr_managment/apis/services"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, store *apis.Store) {
	var (
		user_repo = repositories.NewUserRepository(store)
		user_serv = services.NewUserService(user_repo)
		user_cont = controllers.NewUserController(user_serv)
	)

	user_auth := router.Group("/api")
	{
		user_auth.POST("/register-user", user_cont.CreateUser)
	}
}
