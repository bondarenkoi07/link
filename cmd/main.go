package main

import (
	"log"
	"net/http"
	"time"

	"github.com/bondarenki07/link/internal/domain"
	"github.com/bondarenki07/link/internal/endpoints"
	"github.com/bondarenki07/link/internal/service"
	"github.com/bondarenki07/link/pkg/redis"
	"github.com/bondarenki07/link/pkg/shorter"
	"github.com/bondarenki07/link/pkg/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := domain.NewConfig()
	rs := redis.NewStorage(cfg.RedisHost, cfg.RedisPass)
	cache := storage.NewMemoryCache()
	short := shorter.NewShorter(cfg.BlockSize, cfg.BlockCount, cfg.Seed)

	ss := service.NewShortLink(cache, rs, short)
	es := endpoints.NewShortLink(ss)

	sl := service.NewLoader(cache, rs)
	er := endpoints.NewRedirect(sl)

	r := gin.Default()

	r.POST("", es.Put)
	r.GET(":short", er.Redirect)

	server := &http.Server{
		Addr:           cfg.Port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
