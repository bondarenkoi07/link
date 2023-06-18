package endpoints

import (
	"context"
	"net/http"
	"strings"

	"github.com/bondarenki07/link/internal/domain"
	"github.com/gin-gonic/gin"
)

type Put interface {
	ShortLink(ctx context.Context, link domain.Link) (string, error)
}

type ShortLink struct {
	srv Put
}

func NewShortLink(srv Put) *ShortLink {
	return &ShortLink{srv: srv}
}

func (e ShortLink) Put(ctx *gin.Context) {
	var link domain.Link

	if err := ctx.ShouldBindJSON(&link); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	put, err := e.srv.ShortLink(ctx, link)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	short := domain.Short{Link: strings.Join([]string{ctx.Request.Host, put}, "/")}

	ctx.JSON(http.StatusCreated, short)
}
