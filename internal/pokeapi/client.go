package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	http_client http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		http_client: http.Client{
			Timeout: timeout,
		},
	}
}
