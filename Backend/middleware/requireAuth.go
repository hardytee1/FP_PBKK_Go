package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hardytee1/FP_PBKK_Go/Backend/initializers"
	"github.com/hardytee1/FP_PBKK_Go/Backend/models"
)

func RequireAuth(c *gin.Context) {
	//Get cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode validae it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		//check the expire
		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		userIDStr := claims["sub"].(string)
        userID, err := uuid.Parse(userIDStr)
        if err != nil {
            fmt.Println("Error parsing user ID:", err)
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        var user models.User
        initializers.DB.First(&user, userID)
        if user.ID == uuid.Nil.String() {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

		//attach
		c.Set("user", user)

		//continue

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
