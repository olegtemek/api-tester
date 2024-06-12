package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/olegtemek/api-tester/internal/config"
)

type rpc interface {
	DoRequest(req *http.Request) (err error)
}

type Usecase struct {
	cfg *config.Config
	rpc rpc
}

func New(cfg *config.Config, rpc rpc) *Usecase {
	return &Usecase{
		cfg: cfg,
		rpc: rpc,
	}
}

func (u *Usecase) Start(ctx context.Context) (goodCount int, err error) {
	goodChan := make(chan struct{})
	defer close(goodChan)

	httpReq, err := http.NewRequest(u.cfg.Method, u.cfg.Url, nil)
	if err != nil {
		return
	}

	for wI := 0; wI < u.cfg.WorkerCount; wI++ {
		fmt.Printf("WORKER %v is ready\n", wI)

		for rI := 0; rI < u.cfg.RequestCount; rI++ {

			go func(wI, rI int) {
				select {
				case <-ctx.Done():
					return
				default:
				}
				err := u.rpc.DoRequest(httpReq)
				if err != nil {
					fmt.Printf("ERROR REQUEST: Worker %v. Request %v. ERROR: %v\n", wI, rI, err)
					return
				}
				goodChan <- struct{}{}
			}(wI, rI)

		}
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("TIMEOUT")
			return
		case <-goodChan:
			goodCount++
		case <-time.After(time.Second * 10):
			fmt.Println("TIMEOUT")
			return
		}
	}
}
