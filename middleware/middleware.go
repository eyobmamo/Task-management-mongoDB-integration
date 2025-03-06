package middleware

import (
    "fmt"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"

    "os"
    "github.com/joho/godotenv"
)

var JwtSecret []byte

func init() {
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }
    JwtSecret = []byte(os.Getenv("JWT_SECRET"))
    if len(JwtSecret) == 0 {
        fmt.Println("JWT_SECRET is not set in the environment variables")
    }
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        // JWT validation logic
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(401, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        authParts := strings.Split(authHeader, " ")
        if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
            c.JSON(401, gin.H{"error": "Invalid authorization header"})
            c.Abort()
            return
        }

        token, err := jwt.Parse(authParts[1], func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }

            return JwtSecret, nil
        })
        if err != nil || !token.Valid {
            c.JSON(401, gin.H{"error": "Invalid JWT"})
            c.Abort()
            return
        }

        // Optionally, you can extract claims and set them in the context
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("claims", claims)
        } else {
            c.JSON(401, gin.H{"error": "Invalid JWT claims"})
            c.Abort()
            return
        }

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            userID := claims["user_id"].(string)
            c.Set("userID", userID)
        } else {
            c.JSON(401, gin.H{"error": "Invalid JWT claims"})
            c.Abort()
            return
        }

        c.Next()
    }
}