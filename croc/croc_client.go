package croc

import (
	"net/http"
	"time"
)
type Config struct {
	client *http.Client
	api_url string
	project string
	access_key string
	secret_key string	
}


func newCrocClient(api_url, access_key, secret_key, project string) *Config{
	return &Config{
		client : &http.Client{
			Timeout: time.Second * 10,
		},
		api_url: api_url,
		project: project,
		access_key: access_key,
		secret_key: secret_key,
	}
}
