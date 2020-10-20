package api

import (
	"net/http"

	"github.com/chnejohnson/shortener/domain"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type UserURLHandler struct {
	domain.UserURLService
}

func NewUserURLHandler(app *echo.Group, us domain.UserURLService) {
	h := &UserURLHandler{us}
	app.GET("/urls", h.fetchAll)
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
