package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
)

// DeleteBlog deletes a blog by its ID
func DeleteBlog(c *gin.Context) {
	// Extract blog ID from URL parameters
	blogID := c.Param("id")

	// Get the authenticated user from the context
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

	// Find the blog by ID and UserID
	var blog models.Blog
	result := initializers.DB.First(&blog, "id = ? AND user_id = ?", blogID, usr.ID)
	if result.Error != nil {
		utils.RespondError(c, http.StatusNotFound, "Blog not found", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	// Delete the blog from the database
	result = initializers.DB.Delete(&blog)
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete blog", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	// Respond with a success message
	utils.RespondSuccess(c, nil, "Blog deleted successfully")
}
