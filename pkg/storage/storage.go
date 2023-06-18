package storage

import (
	"errors"
	"sync"
)

var (
	ErrSameShort         = errors.New("link hash collision")
	ErrNotFound          = errors.New("link not found")
	ErrInconvertibleType = errors.New("could not convert to string")
)

type MemoryCache struct {
	storage *sync.Map
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{storage: new(sync.Map)}
}

func (s MemoryCache) Put(short string, link string) error {
	storedLink, shortExists := s.storage.LoadOrStore(short, link)
	if !shortExists {
		return nil
	}

	if storedLink == link {
		return nil
	}

	return ErrSameShort
}

func (s MemoryCache) Load(shortURI string) (string, error) {
	var (
		linkRaw any
		ok      bool
	)

	if linkRaw, ok = s.storage.Load(shortURI); !ok {
		return "", ErrNotFound
	}

	switch v := linkRaw.(type) {
	case string:
		return v, nil
	case []rune:
		return string(v), nil
	default:
		return "", ErrInconvertibleType
	}
}
