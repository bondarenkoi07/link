package service

import (
	"context"
	"log"

	"github.com/bondarenki07/link/internal/domain"
)

type Shorter interface {
	Short(url string) string
}

type ShortCache interface {
	Put(key string, link string) error
}

type StorageShort interface {
	Put(ctx context.Context, key string, link string) error
}

type ShortLink struct {
	cache   ShortCache
	storage StorageShort
	shorter Shorter
}

func NewShortLink(cache ShortCache, storage StorageShort, shorter Shorter) *ShortLink {
	return &ShortLink{cache: cache, storage: storage, shorter: shorter}
}

func (s ShortLink) ShortLink(ctx context.Context, link domain.Link) (string, error) {
	short := s.shorter.Short(link.Link)

	log.Println(short)

	err := s.cache.Put(short, link.Link)
	if err != nil {
		return "", err
	}

	err = s.storage.Put(ctx, short, link.Link)
	if err != nil {
		return "", err
	}

	return short, nil
}
