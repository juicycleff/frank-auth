// Code generated by goa v3.20.0, DO NOT EDIT.
//
// health client HTTP transport
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the health service endpoint HTTP clients.
type Client struct {
	// Check Doer is the HTTP client used to make requests to the check endpoint.
	CheckDoer goahttp.Doer

	// Ready Doer is the HTTP client used to make requests to the ready endpoint.
	ReadyDoer goahttp.Doer

	// Version Doer is the HTTP client used to make requests to the version
	// endpoint.
	VersionDoer goahttp.Doer

	// Metrics Doer is the HTTP client used to make requests to the metrics
	// endpoint.
	MetricsDoer goahttp.Doer

	// Debug Doer is the HTTP client used to make requests to the debug endpoint.
	DebugDoer goahttp.Doer

	// CORS Doer is the HTTP client used to make requests to the  endpoint.
	CORSDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the health service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		CheckDoer:           doer,
		ReadyDoer:           doer,
		VersionDoer:         doer,
		MetricsDoer:         doer,
		DebugDoer:           doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Check returns an endpoint that makes HTTP requests to the health service
// check server.
func (c *Client) Check() goa.Endpoint {
	var (
		decodeResponse = DecodeCheckResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildCheckRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CheckDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("health", "check", err)
		}
		return decodeResponse(resp)
	}
}

// Ready returns an endpoint that makes HTTP requests to the health service
// ready server.
func (c *Client) Ready() goa.Endpoint {
	var (
		decodeResponse = DecodeReadyResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildReadyRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ReadyDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("health", "ready", err)
		}
		return decodeResponse(resp)
	}
}

// Version returns an endpoint that makes HTTP requests to the health service
// version server.
func (c *Client) Version() goa.Endpoint {
	var (
		decodeResponse = DecodeVersionResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildVersionRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.VersionDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("health", "version", err)
		}
		return decodeResponse(resp)
	}
}

// Metrics returns an endpoint that makes HTTP requests to the health service
// metrics server.
func (c *Client) Metrics() goa.Endpoint {
	var (
		decodeResponse = DecodeMetricsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildMetricsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MetricsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("health", "metrics", err)
		}
		return decodeResponse(resp)
	}
}

// Debug returns an endpoint that makes HTTP requests to the health service
// debug server.
func (c *Client) Debug() goa.Endpoint {
	var (
		decodeResponse = DecodeDebugResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDebugRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DebugDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("health", "debug", err)
		}
		return decodeResponse(resp)
	}
}
