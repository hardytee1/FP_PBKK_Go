package routers

import (
	"github.com/hardytee1/FP_PBKK_Go/Backend/controllers/blog"
	"github.com/hardytee1/FP_PBKK_Go/Backend/middleware"

	"github.com/gin-gonic/gin"
)

func BlogRouter(router *gin.Engine) {
	blogRoutes := router.Group("/api/blog")

	blogRoutes.POST("/blog", middleware.RequireAuth, controllers.CreateBlog) 
	blogRoutes.GET("/blogs", middleware.RequireAuth, controllers.GetAllBlog)
	blogRoutes.GET("/blog", middleware.RequireAuth, controllers.GetUserBlogs)
	blogRoutes.PUT("/blog/:id", middleware.RequireAuth, controllers.UpdateBlog)
}