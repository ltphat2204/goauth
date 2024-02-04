package middlewares

import (
	"github.com/gin-gonic/gin"     // request
	"github.com/golang-jwt/jwt/v5" // parse json web token
	"net/http"                     // http status
	"os"                           // load .env
	"strings"                      // upper and compare string
	"time"                         // get now
)

func AuthMiddleware(c *gin.Context) {
	// Get Authorization token
	tokenString := c.GetHeader("Authorization")

	// No Auth
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No Token found!"})
		c.Abort()
		return
	}

	// Get Authorization method
	method := tokenString[:6]
	tokenString = tokenString[7:]

	// Validate token
	if strings.Compare(strings.ToUpper(method), "BEARER") != 0 || tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token!"})
		c.Abort()
		return
	}

	// Parse token from secret key
	secretKey := []byte(os.Getenv("SECRET_KEY"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// Validate token
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token!"})
		c.Abort()
		return
	}

	// Claim token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Token can not claim!"})
		c.Abort()
		return
	}

	// Get username
	username, exists := claims["username"].(string)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Token can not claim!"})
		c.Abort()
		return
	}

	// Token here has passed the verify token process
	// so no need to catch error anymore
	expire, _ := claims.GetExpirationTime()

	// Token expired
	if expire.Time.Unix() < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Expired token!"})
		c.Abort()
		return
	}

	// Pass username to next handler
	c.Set("username", username)
	c.Next()
}
