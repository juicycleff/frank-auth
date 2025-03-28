// Code generated by goa v3.20.0, DO NOT EDIT.
//
// mfa client HTTP transport
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

// Client lists the mfa service endpoint HTTP clients.
type Client struct {
	// Enroll Doer is the HTTP client used to make requests to the enroll endpoint.
	EnrollDoer goahttp.Doer

	// Verify Doer is the HTTP client used to make requests to the verify endpoint.
	VerifyDoer goahttp.Doer

	// Unenroll Doer is the HTTP client used to make requests to the unenroll
	// endpoint.
	UnenrollDoer goahttp.Doer

	// Methods Doer is the HTTP client used to make requests to the methods
	// endpoint.
	MethodsDoer goahttp.Doer

	// SendCode Doer is the HTTP client used to make requests to the send_code
	// endpoint.
	SendCodeDoer goahttp.Doer

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

// NewClient instantiates HTTP clients for all the mfa service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		EnrollDoer:          doer,
		VerifyDoer:          doer,
		UnenrollDoer:        doer,
		MethodsDoer:         doer,
		SendCodeDoer:        doer,
		CORSDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Enroll returns an endpoint that makes HTTP requests to the mfa service
// enroll server.
func (c *Client) Enroll() goa.Endpoint {
	var (
		encodeRequest  = EncodeEnrollRequest(c.encoder)
		decodeResponse = DecodeEnrollResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildEnrollRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.EnrollDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mfa", "enroll", err)
		}
		return decodeResponse(resp)
	}
}

// Verify returns an endpoint that makes HTTP requests to the mfa service
// verify server.
func (c *Client) Verify() goa.Endpoint {
	var (
		encodeRequest  = EncodeVerifyRequest(c.encoder)
		decodeResponse = DecodeVerifyResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildVerifyRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.VerifyDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mfa", "verify", err)
		}
		return decodeResponse(resp)
	}
}

// Unenroll returns an endpoint that makes HTTP requests to the mfa service
// unenroll server.
func (c *Client) Unenroll() goa.Endpoint {
	var (
		encodeRequest  = EncodeUnenrollRequest(c.encoder)
		decodeResponse = DecodeUnenrollResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildUnenrollRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UnenrollDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mfa", "unenroll", err)
		}
		return decodeResponse(resp)
	}
}

// Methods returns an endpoint that makes HTTP requests to the mfa service
// methods server.
func (c *Client) Methods() goa.Endpoint {
	var (
		encodeRequest  = EncodeMethodsRequest(c.encoder)
		decodeResponse = DecodeMethodsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildMethodsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.MethodsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mfa", "methods", err)
		}
		return decodeResponse(resp)
	}
}

// SendCode returns an endpoint that makes HTTP requests to the mfa service
// send_code server.
func (c *Client) SendCode() goa.Endpoint {
	var (
		encodeRequest  = EncodeSendCodeRequest(c.encoder)
		decodeResponse = DecodeSendCodeResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildSendCodeRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.SendCodeDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mfa", "send_code", err)
		}
		return decodeResponse(resp)
	}
}
