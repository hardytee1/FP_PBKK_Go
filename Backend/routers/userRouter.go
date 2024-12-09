package routers

import (
	"github.com/hardytee1/FP_PBKK_Go/Backend/controllers/authenticate"
	"github.com/hardytee1/FP_PBKK_Go/Backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine) {
	userRoutes := router.Group("/api/user")

	// userRoutes.GET("/validate",middleware.RequireAuth, middleware.RequireRole(models.RoleAdmin), controllers.Validate)
	// contoh buat yg akses cuman boleh role tertentu

	userRoutes.POST("/register", controllers.Register)
	userRoutes.POST("/login", controllers.Login)
	userRoutes.POST("/logout", controllers.Logout)
	userRoutes.GET("/me", middleware.RequireAuth, controllers.Me)
	userRoutes.DELETE("/delete", middleware.RequireAuth, controllers.DeleteCurrentUser)
	userRoutes.PUT("/update", middleware.RequireAuth, controllers.Update)
}