// Code generated by goa v3.20.0, DO NOT EDIT.
//
// sso client
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package sso

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "sso" service client.
type Client struct {
	ListProvidersEndpoint          goa.Endpoint
	ProviderAuthEndpoint           goa.Endpoint
	ProviderCallbackEndpoint       goa.Endpoint
	ListIdentityProvidersEndpoint  goa.Endpoint
	CreateIdentityProviderEndpoint goa.Endpoint
	GetIdentityProviderEndpoint    goa.Endpoint
	UpdateIdentityProviderEndpoint goa.Endpoint
	DeleteIdentityProviderEndpoint goa.Endpoint
	SamlMetadataEndpoint           goa.Endpoint
	SamlAcsEndpoint                goa.Endpoint
}

// NewClient initializes a "sso" service client given the endpoints.
func NewClient(listProviders, providerAuth, providerCallback, listIdentityProviders, createIdentityProvider, getIdentityProvider, updateIdentityProvider, deleteIdentityProvider, samlMetadata, samlAcs goa.Endpoint) *Client {
	return &Client{
		ListProvidersEndpoint:          listProviders,
		ProviderAuthEndpoint:           providerAuth,
		ProviderCallbackEndpoint:       providerCallback,
		ListIdentityProvidersEndpoint:  listIdentityProviders,
		CreateIdentityProviderEndpoint: createIdentityProvider,
		GetIdentityProviderEndpoint:    getIdentityProvider,
		UpdateIdentityProviderEndpoint: updateIdentityProvider,
		DeleteIdentityProviderEndpoint: deleteIdentityProvider,
		SamlMetadataEndpoint:           samlMetadata,
		SamlAcsEndpoint:                samlAcs,
	}
}

// ListProviders calls the "list_providers" endpoint of the "sso" service.
// ListProviders may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
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

// ProviderAuth calls the "provider_auth" endpoint of the "sso" service.
// ProviderAuth may return the following errors:
//   - "not_found" (type *NotFoundError): Provider not found
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) ProviderAuth(ctx context.Context, p *ProviderAuthPayload) (err error) {
	_, err = c.ProviderAuthEndpoint(ctx, p)
	return
}

// ProviderCallback calls the "provider_callback" endpoint of the "sso" service.
// ProviderCallback may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
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

// ListIdentityProviders calls the "list_identity_providers" endpoint of the
// "sso" service.
// ListIdentityProviders may return the following errors:
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "bad_request" (type *BadRequestError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) ListIdentityProviders(ctx context.Context, p *ListIdentityProvidersPayload) (res *ListIdentityProvidersResult, err error) {
	var ires any
	ires, err = c.ListIdentityProvidersEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*ListIdentityProvidersResult), nil
}

// CreateIdentityProvider calls the "create_identity_provider" endpoint of the
// "sso" service.
// CreateIdentityProvider may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "not_found" (type *NotFoundError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) CreateIdentityProvider(ctx context.Context, p *CreateIdentityProviderPayload) (res *IdentityProviderResponse, err error) {
	var ires any
	ires, err = c.CreateIdentityProviderEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*IdentityProviderResponse), nil
}

// GetIdentityProvider calls the "get_identity_provider" endpoint of the "sso"
// service.
// GetIdentityProvider may return the following errors:
//   - "not_found" (type *goa.ServiceError)
//   - "unauthorized" (type *goa.ServiceError)
//   - "forbidden" (type *goa.ServiceError)
//   - "bad_request" (type *BadRequestError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) GetIdentityProvider(ctx context.Context, p *GetIdentityProviderPayload) (res *IdentityProviderResponse, err error) {
	var ires any
	ires, err = c.GetIdentityProviderEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*IdentityProviderResponse), nil
}

// UpdateIdentityProvider calls the "update_identity_provider" endpoint of the
// "sso" service.
// UpdateIdentityProvider may return the following errors:
//   - "bad_request" (type *goa.ServiceError)
//   - "not_found" (type *goa.ServiceError)
//   - "unauthorized" (type *goa.ServiceError)
//   - "forbidden" (type *goa.ServiceError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) UpdateIdentityProvider(ctx context.Context, p *UpdateIdentityProviderPayload) (res *IdentityProviderResponse, err error) {
	var ires any
	ires, err = c.UpdateIdentityProviderEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*IdentityProviderResponse), nil
}

// DeleteIdentityProvider calls the "delete_identity_provider" endpoint of the
// "sso" service.
// DeleteIdentityProvider may return the following errors:
//   - "not_found" (type *NotFoundError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "bad_request" (type *BadRequestError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) DeleteIdentityProvider(ctx context.Context, p *DeleteIdentityProviderPayload) (err error) {
	_, err = c.DeleteIdentityProviderEndpoint(ctx, p)
	return
}

// SamlMetadata calls the "saml_metadata" endpoint of the "sso" service.
// SamlMetadata may return the following errors:
//   - "not_found" (type *NotFoundError)
//   - "bad_request" (type *BadRequestError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) SamlMetadata(ctx context.Context, p *SamlMetadataPayload) (res *SamlMetadataResult, err error) {
	var ires any
	ires, err = c.SamlMetadataEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*SamlMetadataResult), nil
}

// SamlAcs calls the "saml_acs" endpoint of the "sso" service.
// SamlAcs may return the following errors:
//   - "bad_request" (type *BadRequestError)
//   - "not_found" (type *NotFoundError)
//   - "unauthorized" (type *UnauthorizedError)
//   - "forbidden" (type *ForbiddenError)
//   - "conflict" (type *ConflictError)
//   - "internal_error" (type *InternalServerError)
//   - error: internal error
func (c *Client) SamlAcs(ctx context.Context, p *SamlAcsPayload) (res string, err error) {
	var ires any
	ires, err = c.SamlAcsEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(string), nil
}
