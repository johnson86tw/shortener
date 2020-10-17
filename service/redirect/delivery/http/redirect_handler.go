package http

import (
	"net/http"

	"github.com/chnejohnson/shortener/domain"
	"github.com/gin-gonic/gin"
)

type redirectHandler struct {
	service domain.RedirectService
}

// NewRedirectHandler ...
func NewRedirectHandler(g *gin.Engine, service domain.RedirectService) {
	handler := &redirectHandler{service}

	g.GET("/:code", handler.find)
	g.POST("/", handler.store)
}

func (r *redirectHandler) find(c *gin.Context) {
	code := c.Params.ByName("code")

	redirect, err := r.service.Find(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "取得重導連結失敗",
		})
		return
	}

	c.Redirect(http.StatusFound, redirect.URL)
}

func (r *redirectHandler) store(c *gin.Context) {
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

	err = r.service.Store(redirect)
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
