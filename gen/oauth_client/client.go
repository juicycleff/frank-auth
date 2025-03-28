// Code generated by goa v3.20.0, DO NOT EDIT.
//
// oauth_client client
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package oauthclient

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "oauth_client" service client.
type Client struct {
	ListProvidersEndpoint    goa.Endpoint
	ProviderAuthEndpoint     goa.Endpoint
	ProviderCallbackEndpoint goa.Endpoint
}

// NewClient initializes a "oauth_client" service client given the endpoints.
func NewClient(listProviders, providerAuth, providerCallback goa.Endpoint) *Client {
	return &Client{
		ListProvidersEndpoint:    listProviders,
		ProviderAuthEndpoint:     providerAuth,
		ProviderCallbackEndpoint: providerCallback,
	}
}

// ListProviders calls the "list_providers" endpoint of the "oauth_client"
// service.
// ListProviders may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) ListProviders(ctx context.Context, p *ListProvidersPayload) (res *ListProvidersResult, err error) {
	var ires any
	ires, err = c.ListProvidersEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ListProvidersResult), nil
}

// ProviderAuth calls the "provider_auth" endpoint of the "oauth_client"
// service.
// ProviderAuth may return the following errors:
//   - "not_found" (type *NotFoundError): Provider not found
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) ProviderAuth(ctx context.Context, p *ProviderAuthPayload) (err error) {
	_, err = c.ProviderAuthEndpoint(ctx, p)
	return
}

// ProviderCallback calls the "provider_callback" endpoint of the
// "oauth_client" service.
// ProviderCallback may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) ProviderCallback(ctx context.Context, p *ProviderCallbackPayload) (res *ProviderCallbackResult, err error) {
	var ires any
	ires, err = c.ProviderCallbackEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ProviderCallbackResult), nil
}
