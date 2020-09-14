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

func (r *redirectHandler) find(ctx *gin.Context) {
	code := ctx.Params.ByName("code")

	rdrt, err := r.service.Find(code)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"retCode": "1",
			"retMsg":  "取得重導連結失敗",
			"retData": nil,
		})
		return
	}

	ctx.Redirect(http.StatusFound, rdrt.URL)
}

func (r *redirectHandler) store(ctx *gin.Context) {
	url := ctx.PostForm("url")
	if url == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"retCode": "1",
			"retMsg":  "沒有網址",
			"retData": nil,
		})
		return
	}

	rdrt := &domain.Redirect{}
	rdrt.URL = url

	err := r.service.Store(rdrt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"retCode": "1",
			"retMsg":  "儲存失敗",
			"retData": nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"retCode": "0",
		"retMsg":  "成功",
		"retData": *rdrt,
	})

}
