package api

import (
	"net/http"

	"github.com/chnejohnson/shortener/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// UserURLHandler ...
type UserURLHandler struct {
	domain.UserURLService
}

// NewUserURLHandler ...
func NewUserURLHandler(app *echo.Group, us domain.UserURLService) {
	h := &UserURLHandler{us}
	app.GET("/urls", h.fetchAll)
	app.POST("/url", h.addURL)
}

func (u *UserURLHandler) fetchAll(c echo.Context) error {
	userID := c.Get("userID").(string)
	uuid := uuid.MustParse(userID)

	urls, err := u.FetchAll(uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			"error": "Fail to get all urls",
		})
	}

	return c.JSON(http.StatusOK, Response{
		"data": urls,
	})
}

func (u *UserURLHandler) addURL(c echo.Context) error {
	var body struct {
		URL string `json:"url" binding:"required"`
	}

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			"error": err.Error(),
		})
	}

	userID := c.Get("userID").(string)
	uuid := uuid.MustParse(userID)

	uu := &domain.UserURL{}
	uu.URL = body.URL
	uu.UserID = uuid

	err = u.AddURL(uu)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			"error": err,
		})
	}

	return c.JSON(http.StatusOK, Response{
		"message": "success",
		"data":    *uu,
	})

}
