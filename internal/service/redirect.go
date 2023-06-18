package service

import (
	"errors"
	"log"

	"github.com/bondarenki07/link/pkg/storage"
	"golang.org/x/net/context"
)

type Cache interface {
	Load(shortURI string) (string, error)
	Put(key string, link string) error
}

type Storage interface {
	Get(ctx context.Context, key string) (string, error)
}

type Loader struct {
	cache   Cache
	storage Storage
}

func NewLoader(cache Cache, storage Storage) *Loader {
	return &Loader{cache: cache, storage: storage}
}

func (s Loader) LoadURI(ctx context.Context, uri string) (string, error) {
	link, err := s.cache.Load(uri)
	if err == nil {
		return link, nil
	}
	if !errors.Is(err, storage.ErrNotFound) {
		return "", err
	}

	link, err = s.storage.Get(ctx, uri)
	if err != nil {
		return "", err
	}

	err = s.cache.Put(uri, link)
	if err != nil {
		log.Printf("could not cache uri %s for link %s", uri, link)
	}

	return link, nil
}
