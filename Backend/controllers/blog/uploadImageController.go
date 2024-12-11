package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
)

func UploadImage(c *gin.Context) {
	// Get the file from the form data
	file, _ := c.FormFile("image")
	if file == nil {
		utils.RespondError(c, http.StatusBadRequest, "No file uploaded", nil)
		return
	}

	// Validate file type (only allow jpg, jpeg, png)
	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	fileExtension := filepath.Ext(file.Filename)
	if !contains(allowedExtensions, fileExtension) {
		utils.RespondError(c, http.StatusBadRequest, "Invalid file type, only JPG, JPEG, PNG are allowed", nil)
		return
	}

	// Create the upload path
	uploadPath := "./uploads/blogs/"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		err := os.MkdirAll(uploadPath, os.ModePerm)
		if err != nil {
			utils.RespondError(c, http.StatusInternalServerError, "Failed to create upload directory", nil)
			return
		}
	}

	// Generate a unique filename for the image
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filePath := filepath.Join(uploadPath, filename)

	// Save the file to the server
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to save file", nil)
		return
	}

	// Return the image URL in the response

	// Optionally, you can store the image URL in the database associated with the blog

	// For now, we'll just return the image URL in the response.

	c.JSON(http.StatusOK, gin.H{"content": filePath})
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
