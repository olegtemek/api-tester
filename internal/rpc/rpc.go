package rpc

import (
	"fmt"
	"net/http"

	"github.com/olegtemek/api-tester/internal/config"
)

type Rpc struct {
	cfg    *config.Config
	client *http.Client
}

func New(cfg *config.Config, client *http.Client) *Rpc {
	return &Rpc{
		cfg:    cfg,
		client: client,
	}
}

func (r *Rpc) DoRequest(req *http.Request) (err error) {
	res, err := r.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("status is not ok Status: %s", res.Status)
		return
	}
	return
}
