package controller

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Claims ...
type Claims struct {
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

// JWT ...
type JWT struct {
	JWTSecret []byte
}

// AuthRequired ...
func (j *JWT) AuthRequired(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	bears := strings.Split(auth, "Bearer ")

	if len(bears) < 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "no bearing",
		})
		c.Abort()
		return
	}

	token := bears[1]

	// parse and validate token for six things:
	// validationErrorMalformed => token is malformed
	// validationErrorUnverifiable => token could not be verified because of signing problems
	// validationErrorSignatureInvalid => signature validation failed
	// validationErrorExpired => exp validation failed
	// validationErrorNotValidYet => nbf validation failed
	// validationErrorIssuedAt => iat validation failed
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(j.JWTSecret), nil
	})

	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		c.Set("account", claims.Account)
		c.Set("role", claims.Role)
		c.Next()
	} else {
		c.Abort()
		return
	}
}
