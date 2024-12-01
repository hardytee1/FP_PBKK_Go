package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
)

func GetUserBlogs(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		utils.RespondError(c, http.StatusUnauthorized, "User not found in context", nil)
		return
	}

	usr, ok := user.(models.User)
	if !ok {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid user type in context", nil)
		return
	}

	var blogs []models.Blog
	result := initializers.DB.Where("user_id = ?", usr.ID).Find(&blogs)
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to retrieve blogs", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	type ResponseData struct {
		ID        uint    `json:"id"`
		Content   string    `json:"content"`
		Caption   string    `json:"caption"`
		UserID    string    `json:"user_id"`
		UserName string    `json:"user_name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	var responseData []ResponseData
	for _, blog := range blogs {
		responseData = append(responseData, ResponseData{
			ID:        blog.ID,
			Content:   blog.Content,
			Caption:   blog.Caption,
			UserID:    blog.UserID,
			UserName:  usr.Name,
			CreatedAt: blog.CreatedAt,
			UpdatedAt: blog.UpdatedAt,
		})
	}

	utils.RespondSuccess(c, responseData, "Blogs retrieved successfully")
}