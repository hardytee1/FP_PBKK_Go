package controllers

import (
	"net/http"
	"os"
	"time"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	//get the email pass
	var body struct {
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	//look up requested user
	var user models.User
	initializers.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == uuid.Nil.String() {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid email or password", gin.H{"username": "Invalid credentials"})
		return
	}

	//compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		utils.RespondError(c, http.StatusUnauthorized, "Invalid email or password", gin.H{"password": "Invalid credentials"})
		return
	}

	//generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to generate token", nil)
		return
	}

	//Send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	type ResponseData struct {
		ID       string    `json:"id"`
		Name string `json:"name"`
		Email    string `json:"email"`
		Token    string `json:"token"`
	}

	responseData :=  ResponseData{
		ID:       user.ID,
		Name : user.Name,
		Email:    user.Email,
		Token:    tokenString,
	}

	utils.RespondSuccess(c, responseData, "Login successful")
}