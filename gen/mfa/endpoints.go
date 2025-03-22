// Code generated by goa v3.20.0, DO NOT EDIT.
//
// mfa endpoints
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package mfa

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "mfa" service endpoints.
type Endpoints struct {
	Enroll   goa.Endpoint
	Verify   goa.Endpoint
	Unenroll goa.Endpoint
	Methods  goa.Endpoint
	SendCode goa.Endpoint
}

// NewEndpoints wraps the methods of the "mfa" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Enroll:   NewEnrollEndpoint(s, a.JWTAuth),
		Verify:   NewVerifyEndpoint(s, a.JWTAuth),
		Unenroll: NewUnenrollEndpoint(s, a.JWTAuth),
		Methods:  NewMethodsEndpoint(s, a.JWTAuth),
		SendCode: NewSendCodeEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "mfa" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Enroll = m(e.Enroll)
	e.Verify = m(e.Verify)
	e.Unenroll = m(e.Unenroll)
	e.Methods = m(e.Methods)
	e.SendCode = m(e.SendCode)
}

// NewEnrollEndpoint returns an endpoint function that calls the method
// "enroll" of service "mfa".
func NewEnrollEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*EnrollPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Enroll(ctx, p)
	}
}

// NewVerifyEndpoint returns an endpoint function that calls the method
// "verify" of service "mfa".
func NewVerifyEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*VerifyPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Verify(ctx, p)
	}
}

// NewUnenrollEndpoint returns an endpoint function that calls the method
// "unenroll" of service "mfa".
func NewUnenrollEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UnenrollPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Unenroll(ctx, p)
	}
}

// NewMethodsEndpoint returns an endpoint function that calls the method
// "methods" of service "mfa".
func NewMethodsEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*MethodsPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.Methods(ctx, p)
	}
}

// NewSendCodeEndpoint returns an endpoint function that calls the method
// "send_code" of service "mfa".
func NewSendCodeEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*SendCodePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:read", "api:write"},
			RequiredScopes: []string{},
		}
		var token string
		if p.JWT != nil {
			token = *p.JWT
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return s.SendCode(ctx, p)
	}
}
