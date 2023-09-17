package api

import (
	"net/http"
	"time"

	"github.com/charlesozo/pokedex/internal/cache"
)
const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	cache cache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client{
  return Client{
	cache: cache.NewCache(cacheInterval),
	httpClient:  http.Client{
		Timeout : timeout,
	},
  }
}


