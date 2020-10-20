package api

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
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
func (j *JWT) AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		bears := strings.Split(auth, "Bearer ")

		if len(bears) < 2 {
			return c.JSON(http.StatusUnauthorized, Response{
				"error": "no bearing",
			})
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
			return c.JSON(http.StatusUnauthorized, Response{
				"error": message,
			})
		}

		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			c.Set("userID", claims.Subject)
			c.Set("role", claims.Role)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, Response{
			"error": "invalid token",
		})
	}
}
