package pokeapi

import (
	"net/http"
	"time"

	"github.com/arunPdl02/pokedexcli/internal/pokecache"
)

type Client struct {
	http_client http.Client
	cache       *pokecache.Cache
}

func NewClient(timeout time.Duration) Client {
	return Client{
		http_client: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(5 * time.Second),
	}
}
