package api

import (
	"net/http"

	"github.com/chnejohnson/shortener/domain"
	"github.com/gin-gonic/gin"
)

// RedirectHandler ...
type RedirectHandler struct {
	domain.RedirectService
}

// NewRedirectHandler ...
func NewRedirectHandler(engine *gin.Engine, rs domain.RedirectService) {
	h := &RedirectHandler{rs}
	engine.GET("/:code", h.find)
	engine.POST("/", h.store)
}

func (r *RedirectHandler) find(c *gin.Context) {
	code := c.Params.ByName("code")

	redirect, err := r.Find(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "取得重導連結失敗",
		})
		return
	}

	c.Redirect(http.StatusFound, redirect.URL)
}

func (r *RedirectHandler) store(c *gin.Context) {
	var body struct {
		URL string `json:"url" binding:"required"`
	}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	redirect := &domain.Redirect{}
	redirect.URL = body.URL

	err = r.Store(redirect)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": *redirect,
	})

}
