package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/olegtemek/api-tester/internal/config"
	"github.com/olegtemek/api-tester/internal/rpc"
	"github.com/olegtemek/api-tester/internal/usecase"
)

func main() {
	cfg := config.New()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	rpc := rpc.New(cfg, &http.Client{})
	srv := usecase.New(cfg, rpc)

	start := time.Now()

	goodCount, err := srv.Start(ctx)
	if err != nil {
		fmt.Printf("ERROR: %v \n", err)
	}

	endTime := time.Now()
	end := endTime.Sub(start)

	fmt.Printf("\n\nTOTAL REQUESTS: %v \nTOTAL TIME FOR ALL REQUESTS %v", goodCount, end)

}
