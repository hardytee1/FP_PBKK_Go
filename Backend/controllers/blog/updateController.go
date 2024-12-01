package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
)

func UpdateBlog(c *gin.Context) {
	blogID := c.Param("id")

	var body struct {
		Content string `json:"content" binding:"required"`
		Caption string `json:"caption" binding:"required"`
	}
	if err := c.Bind(&body); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request body", map[string]interface{}{"error": err.Error()})
		return
	}

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

	var blog models.Blog
	result := initializers.DB.First(&blog, "id = ? AND user_id = ?", blogID, usr.ID)
	if result.Error != nil {
		utils.RespondError(c, http.StatusNotFound, "Blog not found or does not belong to the user", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	blog.Content = body.Content
	blog.Caption = body.Caption

	if err := initializers.DB.Save(&blog).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update blog", map[string]interface{}{"error": err.Error()})
		return
	}

	utils.RespondSuccess(c, blog, "Blog updated successfully")
}
