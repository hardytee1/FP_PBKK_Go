package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
)

func Me(c *gin.Context) {
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

	tokenString, err := c.Cookie("Authorization")
	
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}


	type ResponseData struct {
		ID       string    `json:"id"`
		Name string `json:"name"`
		Email    string `json:"email"`
		Picture  string `json:"picture"`
		Token    string `json:"token"`
	}

	responseData :=  ResponseData{
		ID:       usr.ID,
		Name: usr.Name,
		Email:    usr.Email,
		Picture:  usr.Picture,
		Token:	  tokenString,
	}

	utils.RespondSuccess(c, responseData, "User validated successfully")
}