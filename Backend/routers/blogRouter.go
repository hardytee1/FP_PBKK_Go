package routers

import (
	controllers "github.com/hardytee1/FP_PBKK_Go/Backend/controllers/blog"
	"github.com/hardytee1/FP_PBKK_Go/Backend/middleware"

	"github.com/gin-gonic/gin"
)

func BlogRouter(router *gin.Engine) {
	blogRoutes := router.Group("/api/blog")

	blogRoutes.POST("/blog", middleware.RequireAuth, controllers.CreateBlog)
	blogRoutes.GET("/blogs", controllers.GetAllBlog)
	blogRoutes.GET("/blog", middleware.RequireAuth, controllers.GetUserBlogs)
	blogRoutes.DELETE("/blog/:id", middleware.RequireAuth, controllers.DeleteBlog)
	blogRoutes.POST("/upload", middleware.RequireAuth, controllers.UploadImage)
	blogRoutes.Static("/uploads", "./uploads")
	blogRoutes.PUT("/update/:id", middleware.RequireAuth, controllers.UpdateBlog)
	blogRoutes.POST("/update/:id", middleware.RequireAuth, controllers.UpdateBlog) // Mengizinkan akses ke folder uploads
}
