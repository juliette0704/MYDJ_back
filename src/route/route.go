package route

import (
	"mydj_server/src/middleware.go"
	"mydj_server/src/shot"
	"mydj_server/src/user"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func RoutingUser(router *gin.Engine, db *gorm.DB) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/login", user.UserLoginController)
		userGroup.POST("/register", user.RegisterController)
	}
}

func RoutingShot(router *gin.Engine, db *gorm.DB) {
	shotGroup := router.Group("/shot")
	{
		shotGroup.Use(middleware.JwtAuthMiddleware())

		shotGroup.POST("/add_shot", shot.AddShotController)
	}
}
