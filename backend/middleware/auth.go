package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	// 1. Authorization header lena
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authorization header required",
		})
		return
	}

	// 2. Bearer TOKEN ko split karna
	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid authorization header",
		})
		return
	}

	tokenString := parts[1]

	// 3. JWT secret lena
	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "JWT_SECRET is not configured",
		})
		return
	}

	// 4. JWT verify / parse karna
	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			// Check signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(secret), nil
		},
	)

	// 5. Token invalid hai
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid or expired token",
		})
		return
	}

	// 6. Claims nikalna
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid token claims",
		})
		return
	}

	// 7. user_id nikalna
	userIDFloat, ok := claims["user_id"].(float64)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user_id missing from token",
		})
		return
	}

	userID := uint(userIDFloat)

	// 8. User ID ko Gin context mein save karna
	c.Set("user_id", userID)

	// 9. Next handler ko allow karna
	c.Next()
}
