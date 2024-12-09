package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
)

func DeleteCurrentUser(c *gin.Context) {
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

	//forcefully delete the blogs of the user
	if err := initializers.DB.Unscoped().Where("user_id = ?", usr.ID).Delete(&models.Blog{}).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete user blogs", map[string]interface{}{"error": err.Error()})
		return
	}

	if err := initializers.DB.Delete(&usr).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete user", map[string]interface{}{"error": err.Error()})
		return
	}

	c.SetCookie("Authorization", "", -1, "", "", false, true)

	utils.RespondSuccess(c, nil, "User deleted successfully")
}
