package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/chnejohnson/shortener/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AccountHandler ...
type AccountHandler struct {
	as domain.AccountService
	*JWT
}

// // try
// acc := &domain.Account{}
// acc.Name = "Howard"
// acc.Email = "howard@gmail.com"
// acc.Password = "23"

// err = as.Create(acc)
// if err != nil {
// 	logrus.Error(err)
// } else {
// 	logrus.Info("Succeed to sign up")
// }

// err = as.Login("howard@gmail.com", "23")
// if err != nil {
// 	logrus.Error(err)
// } else {
// 	logrus.Info("Succeed to login")
// }

// NewAccountHandler ...
func NewAccountHandler(r *gin.Engine, as domain.AccountService, j *JWT) {
	h := &AccountHandler{as, j}
	r.POST("/signup", h.Signup)
	r.POST("login", h.Login)
}

// Signup ...
func (h *AccountHandler) Signup(c *gin.Context) {
	var body struct {
		Email    string
		Password string
		Name     string
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	acc := &domain.Account{}
	acc.Email = body.Email
	acc.Password = body.Password
	acc.Name = body.Name

	err = h.as.Create(acc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})

}

// Login ...
func (h *AccountHandler) Login(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// service
	err = h.as.Login(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	// JWT
	now := time.Now()
	jwtID := body.Email + strconv.FormatInt(now.Unix(), 10)
	role := "Member"

	claims := Claims{
		Account: body.Email,
		Role:    role,
		StandardClaims: jwt.StandardClaims{
			Audience:  body.Email,
			ExpiresAt: now.Add(20 * time.Minute).Unix(),
			Id:        jwtID,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			NotBefore: now.Add(10 * time.Second).Unix(),
			Subject:   body.Email,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(h.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
