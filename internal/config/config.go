package config

import (
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/olegtemek/api-tester/internal/utils"
)

type Config struct {
	Url          string
	WorkerCount  int
	RequestCount int
	Timeout      int
	Method       string
}

func New() *Config {
	testedUrl := flag.String("u", "", "tested url")
	method := flag.String("m", "", "method request")
	workerCount := flag.Int("w", 2, "count of workers for send requests")
	requestCount := flag.Int("c", 1000, "request count")
	timeout := flag.Int("t", 10, "timeout")
	flag.Parse()

	if testedUrl == nil || *testedUrl == "" {
		fmt.Println("cannot parse url, please use -u flag")
		os.Exit(1)
	}

	_, err := url.ParseRequestURI(*testedUrl)
	if err != nil {
		fmt.Println("url must be a valid URL")
		os.Exit(1)
	}

	if method == nil || !utils.ValidateMethod(*method) {
		fmt.Println("method not allowed")
		os.Exit(1)
	}

	return &Config{
		Url:          *testedUrl,
		WorkerCount:  *workerCount,
		RequestCount: *requestCount,
		Method:       *method,
		Timeout:      *timeout,
	}
}
