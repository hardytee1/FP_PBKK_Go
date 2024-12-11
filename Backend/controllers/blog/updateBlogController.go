package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
)

func UpdateBlog(c *gin.Context) {
	// Get blog ID from URL parameter
	blogID := c.Param("id")

	// Define request body structure
	var body struct {
		Content *multipart.FileHeader `form:"content"`
		Caption string                `form:"caption" json:"caption"`
	}

	// Bind request body
	if err := c.Bind(&body); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request body", map[string]interface{}{"error": err.Error()})
		return
	}

	// Get user from context
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

	// Find existing blog
	var blog models.Blog
	if err := initializers.DB.Where("id = ? AND user_id = ?", blogID, usr.ID).First(&blog).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Blog not found or unauthorized", nil)
		return
	}

	// Update blog fields if provided

	if body.Caption != "" {
		blog.Caption = body.Caption
	}

	if body.Content == nil && body.Caption != "" {
		blog.Caption = body.Caption
	} else if body.Content != nil {
		// Logika untuk mengganti content
		if blog.Content != "" {
			fullPath := filepath.Join(".", blog.Content)
			if err := os.Remove(fullPath); err != nil {
				utils.RespondError(c, http.StatusInternalServerError, "Failed to delete existing Content", map[string]interface{}{"error": err.Error()})
				return
			}
		}

		uuidStr := usr.ID
		ext := filepath.Ext(body.Content.Filename)
		filename := fmt.Sprintf("%s_%d%s", uuidStr, time.Now().Unix(), ext)
		savePath := "uploads/blog/" + filename

		if err := c.SaveUploadedFile(body.Content, savePath); err != nil {
			utils.RespondError(c, http.StatusInternalServerError, "Could not save Content", map[string]interface{}{"error": err.Error()})
			return
		}

		blog.Content = savePath
	}

	if err := initializers.DB.Save(&blog).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update blog", map[string]interface{}{"error": err.Error()})
		return
	}

	utils.RespondSuccess(c, blog, "Blog updated successfully")
}
