package endpoints

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoadURI interface {
	LoadURI(ctx context.Context, uri string) (string, error)
}

type Redirect struct {
	srv LoadURI
}

func NewRedirect(srv LoadURI) *Redirect {
	return &Redirect{srv: srv}
}

func (e Redirect) Redirect(ctx *gin.Context) {
	short := ctx.Param("short")
	link, err := e.srv.LoadURI(ctx, short)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, link)
}
