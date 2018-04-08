package memory

import (
	"context"
	"time"

	"github.com/zgiber/cache"
)

func NewAuthStore() *AuthStore {
	return &AuthStore{
		c: cache.NewMemCache(cache.MaxBytesLimit(128 * 1024 * 1024)),
	}
}

type AuthStore struct {
	c cache.Cache
}

func (as *AuthStore) ClientIDByToken(ctx context.Context, token string) (string, bool, error) {
	clientID, err := as.c.Fetch(token)
	switch err {
	case nil:
		return string(clientID), true, nil
	case cache.ErrNotFound:
		return "", false, nil
	default:
		return "", false, err
	}
}

func (as *AuthStore) SetAccessToken(ctx context.Context, clientID, token string) error {
	return as.c.Set(clientID, []byte(token), 168*time.Hour)
}

func (as *AuthStore) DeleteAccessTokens(ctx context.Context, clientID string) error {
	return as.c.Delete(clientID)
}
