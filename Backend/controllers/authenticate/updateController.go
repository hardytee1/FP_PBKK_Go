package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
)

func Update(c *gin.Context) {
	var body struct {
		Name    string `json:"name" binding:"required"`
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

	usr.Name = body.Name

	if err := initializers.DB.Save(&usr).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update User", map[string]interface{}{"error": err.Error()})
		return
	}

	utils.RespondSuccess(c, usr, "User updated successfully")
}
