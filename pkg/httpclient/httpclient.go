package httpclient

import (
	"context"
	"net/http"
	"time"
)

const defaultTimeout = 10 * time.Second

type Client struct {
	client *http.Client
}

func New(timeout ...time.Duration) *Client {
	t := defaultTimeout
	if len(timeout) > 0 {
		t = timeout[0]
	}
	return &Client{
		client: &http.Client{Timeout: 10 *t},
	}
}

func (c *Client) Get(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.client.Do(req)
}

func (c *Client) GetWithTimeout(ctx context.Context, url string, timeout time.Duration) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	return c.Get(ctx, url)
}