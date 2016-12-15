package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// InboundMessagePath computes a request path to the inbound action of message.
func InboundMessagePath() string {
	return fmt.Sprintf("/v1/slack/inbound")
}

// InboundMessage makes a request to the inbound action endpoint of the message resource
func (c *Client) InboundMessage(ctx context.Context, path string, payload *SlackMessage) (*http.Response, error) {
	req, err := c.NewInboundMessageRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewInboundMessageRequest create the request corresponding to the inbound action endpoint of the message resource.
func (c *Client) NewInboundMessageRequest(ctx context.Context, path string, payload *SlackMessage) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	return req, nil
}
