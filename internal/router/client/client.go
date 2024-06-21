package client

import (
	"bytes"
	"context"
	"fmt"
	"github.com/whenborsch/loft2rent-collector-agent/config"
	"github.com/whenborsch/loft2rent-collector-agent/pkg/logger"
	"net/http"
)

var globalClient *Client

type Client struct {
	client *http.Client

	log *logger.Logger
}

func New(cfg config.HTTPClientConfig) *Client {
	if globalClient != nil {
		return globalClient
	}

	c := &Client{
		client: &http.Client{
			Timeout: cfg.Timeout,
		},
		log: logger.GetInstance(),
	}

	if globalClient == nil {
		globalClient = c
	}

	return c
}

func (c *Client) DoRequest(ctx context.Context, req *Request) (*http.Response, error) {
	c.log.Debug(fmt.Sprintf("%s %s", req.method, req.url))

	reader := bytes.NewReader(req.body)
	httpReq, err := http.NewRequest(req.method, req.url, reader)

	if err != nil {
		return nil, err
	}

	for k, v := range req.headers {
		httpReq.Header.Set(k, v)
	}

	httpReq = httpReq.WithContext(ctx)

	return c.client.Do(httpReq)
}
