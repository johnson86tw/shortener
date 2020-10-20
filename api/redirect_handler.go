package api

import (
	"net/http"

	"github.com/chnejohnson/shortener/domain"
	"github.com/labstack/echo"
)

// RedirectHandler ...
type RedirectHandler struct {
	domain.RedirectService
}

// NewRedirectHandler ...
func NewRedirectHandler(app *echo.Echo, rs domain.RedirectService) {
	h := &RedirectHandler{rs}
	app.GET("/:code", h.redirect)
	app.POST("/", h.store)
}

func (r *RedirectHandler) redirect(c echo.Context) error {
	code := c.Param("code")

	redirect, err := r.Find(code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			"error": "Fail to get redirect url",
		})
	}

	return c.Redirect(http.StatusFound, redirect.URL)
}

func (r *RedirectHandler) store(c echo.Context) error {
	var body struct {
		URL string `json:"url" binding:"required"`
	}

	err := c.Bind(&body)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			"error": err.Error(),
		})

	}

	redirect := &domain.Redirect{}
	redirect.URL = body.URL

	err = r.Store(redirect)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Response{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, Response{
		"data": *redirect,
	})

}
