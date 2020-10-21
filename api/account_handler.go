package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/chnejohnson/shortener/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

// AccountHandler ...
type AccountHandler struct {
	domain.AccountService
	*JWT
}

// NewAccountHandler ...
func NewAccountHandler(app *echo.Echo, as domain.AccountService, j *JWT) {
	h := &AccountHandler{as, j}
	app.POST("/signup", h.signup)
	app.POST("/login", h.login)
}

// Signup ...
func (h *AccountHandler) signup(c echo.Context) error {
	var body struct {
		Name     string `json:"name" form:"name" query:"name"`
		Email    string `json:"email" form:"email" query:"email"`
		Password string `json:"password" form:"password" query:"password"`
	}

	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			"error": err.Error(),
		})
	}

	acc := &domain.Account{}
	acc.Email = body.Email
	acc.Password = body.Password
	acc.Name = body.Name

	err := h.Create(acc)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			"error": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, Response{
		"message": "Success",
	})

}

// Login ...
func (h *AccountHandler) login(c echo.Context) error {
	var body struct {
		Email    string `json:"email" form:"email" query:"email"`
		Password string `json:"password" form:"password" query:"password"`
	}

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			"error": err.Error(),
		})
	}

	logrus.WithFields(logrus.Fields{
		"body":     body.Email,
		"password": body.Password,
	}).Info("Login Request body")

	// service
	uuid, err := h.Login(body.Email, body.Password)
	if err != nil {
		// 這裡要想辦法去區分是sql取資料的錯誤，還是service層的錯誤，還是真的密碼沒過
		return c.JSON(http.StatusUnauthorized, Response{
			"message": err.Error(),
		})

	}

	// JWT
	now := time.Now()
	jwtID := strconv.FormatInt(now.Unix(), 10)
	role := "Member"

	claims := Claims{
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(20 * time.Minute).Unix(),
			Id:        jwtID,
			IssuedAt:  now.Unix(),
			Issuer:    "johnson chen",
			NotBefore: now.Add(10 * time.Second).Unix(),
			Subject:   uuid.String(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(h.JWTSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			"error": err.Error(),
		})

	}

	return c.JSON(http.StatusOK, Response{
		"token": token,
	})

}
