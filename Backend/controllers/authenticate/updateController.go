package controllers

import (
	"net/http"
	"mime/multipart"
	"os"
	"path/filepath"
	"fmt"
	"time"


	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
)

func Update(c *gin.Context) {
	var body struct {
		Name    string `json:"name" form:"name"`
		Picture *multipart.FileHeader `form:"picture"`
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

	if body.Name != "" {
		usr.Name = body.Name
	}

	if body.Picture != nil {

		if usr.Picture != "" {
			if err := os.Remove(usr.Picture); err != nil {
				utils.RespondError(c, http.StatusInternalServerError, "Failed to delete existing file", map[string]interface{}{"error": err.Error()})
				return
			}
		}

		uuidStr := usr.ID
		ext := filepath.Ext(body.Picture.Filename) 
		filename := fmt.Sprintf("%s_%d%s", uuidStr, time.Now().Unix(), ext)
		savePath := "uploads/user/" + filename

		if err := c.SaveUploadedFile(body.Picture, savePath); err != nil {
			utils.RespondError(c, http.StatusInternalServerError, "Could not save file", map[string]interface{}{"error": err.Error()})
			return
		}

		usr.Picture = savePath
	}

	if err := initializers.DB.Save(&usr).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update User", map[string]interface{}{"error": err.Error()})
		return
	}

	utils.RespondSuccess(c, usr, "User updated successfully")
}
