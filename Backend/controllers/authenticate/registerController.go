package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
	"github.com/hardytee1/FP_PBKK_Go/Backend/utils"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Failed to register", map[string]interface{}{"error": err.Error()})
		return
	}

	//check if email already exists
	var user_check models.User
	result_check_1 := initializers.DB.Where("email = ?", body.Email).First(&user_check)
	if result_check_1.Error == nil {
		utils.RespondError(c, http.StatusBadRequest, "Email already exists", nil)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to hash password", map[string]interface{}{"error": err.Error()})
		return
	}

	user := models.User{
		Name:     body.Name,
		Password: string(hash),
		Email:    body.Email,
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create user", map[string]interface{}{"error": result.Error.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to generate JWT token", map[string]interface{}{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	utils.RespondSuccess(c, nil, "User registered successfully")
}