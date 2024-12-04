package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
)

func CreateBlog(c *gin.Context) {
	var body struct {
		Content string `json:"content" binding:"required"`
		Caption string `json:"caption" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Failed to create blog", map[string]interface{}{"error": err.Error()})
		return
	}

	user, exists := c.Get("user")
	if !exists {
		utils.RespondError(c, http.StatusUnauthorized, "User not found in context", nil)
		return
	}

	usr, ok := user.(models.User) // Type assertion
	if !ok {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid user type in context", nil)
		return
	}

	blog := models.Blog{
		Content:     body.Content,
		Caption:    body.Caption,
		UserID:   usr.ID,
	}

	result := initializers.DB.Create(&blog)
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create blog", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	//return the data of the created blog
	type ResponseData struct {
		ID       uint    `json:"id"`
		Content string `json:"content"`
		Caption    string `json:"caption"`
		UserID    string `json:"user_id"`
		CreatedAt time.Time `json:"created_at"`
		UpdateAt  time.Time `json:"updated_at"`
	}

	responseData :=  ResponseData{
		ID:       blog.ID,
		Content: blog.Content,
		Caption:    blog.Caption,
		UserID:    blog.UserID,
		CreatedAt: blog.CreatedAt,
		UpdateAt:  blog.UpdatedAt,
	}

	utils.RespondSuccess(c, responseData, "Blog created successfully")
}