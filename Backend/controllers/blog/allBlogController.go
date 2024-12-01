package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
)

func GetAllBlog(c *gin.Context) {

	var blogs []models.Blog
	result := initializers.DB.Find(&blogs)
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
		var usr models.User
		result := initializers.DB.First(&usr, "id = ?", blog.UserID)
		if result.Error != nil {
			utils.RespondError(c, http.StatusInternalServerError, "Failed to retrieve user", map[string]interface{}{"error": result.Error.Error()})
			return
		}
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