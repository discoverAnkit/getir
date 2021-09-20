package repository

import (
	"context"
	"github.com/patrickmn/go-cache"
	"log"
)

type InMemoryClient interface {
	GetValue(ctx context.Context, key string) string
	SetKeyValue(ctx context.Context, key, value string) error
}

type inMemoryClient struct{
	cacheClient   *cache.Cache
}

func NewInMemoryClient (cacheClient *cache.Cache) InMemoryClient {
	return &inMemoryClient{
		cacheClient:cacheClient,
	}
}

func (c *inMemoryClient) GetValue (ctx context.Context, key string) string {

	cached := c.cacheClient
	value ,found := cached.Get(key)
	if found {
		return value.(string)
	}
	log.Printf("Key : %s not found \n",key)
	return ""
}

func (c *inMemoryClient) SetKeyValue(ctx context.Context, key,value string) error {

	cached := c.cacheClient
	cached.Delete(key) //coz this library wont override
	return cached.Add(key,value,cache.NoExpiration)
}