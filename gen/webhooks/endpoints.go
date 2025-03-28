// Code generated by goa v3.20.0, DO NOT EDIT.
//
// webhooks endpoints
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package webhooks

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "webhooks" service endpoints.
type Endpoints struct {
	List         goa.Endpoint
	Create       goa.Endpoint
	Get          goa.Endpoint
	Update       goa.Endpoint
	Delete       goa.Endpoint
	TriggerEvent goa.Endpoint
	ListEvents   goa.Endpoint
	ReplayEvent  goa.Endpoint
	Receive      goa.Endpoint
}

// NewEndpoints wraps the methods of the "webhooks" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		List:         NewListEndpoint(s, a.JWTAuth),
		Create:       NewCreateEndpoint(s, a.JWTAuth),
		Get:          NewGetEndpoint(s, a.JWTAuth),
		Update:       NewUpdateEndpoint(s, a.JWTAuth),
		Delete:       NewDeleteEndpoint(s, a.JWTAuth),
		TriggerEvent: NewTriggerEventEndpoint(s, a.JWTAuth),
		ListEvents:   NewListEventsEndpoint(s, a.JWTAuth),
		ReplayEvent:  NewReplayEventEndpoint(s, a.JWTAuth),
		Receive:      NewReceiveEndpoint(s),
	}
}

// Use applies the given middleware to all the "webhooks" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Create = m(e.Create)
	e.Get = m(e.Get)
	e.Update = m(e.Update)
	e.Delete = m(e.Delete)
	e.TriggerEvent = m(e.TriggerEvent)
	e.ListEvents = m(e.ListEvents)
	e.ReplayEvent = m(e.ReplayEvent)
	e.Receive = m(e.Receive)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "webhooks".
func NewListEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListPayload)
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
		return s.List(ctx, p)
	}
}

// NewCreateEndpoint returns an endpoint function that calls the method
// "create" of service "webhooks".
func NewCreateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreatePayload)
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
		return s.Create(ctx, p)
	}
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "webhooks".
func NewGetEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetPayload)
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
		return s.Get(ctx, p)
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "webhooks".
func NewUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdatePayload)
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
		return s.Update(ctx, p)
	}
}

// NewDeleteEndpoint returns an endpoint function that calls the method
// "delete" of service "webhooks".
func NewDeleteEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeletePayload)
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
		return nil, s.Delete(ctx, p)
	}
}

// NewTriggerEventEndpoint returns an endpoint function that calls the method
// "trigger_event" of service "webhooks".
func NewTriggerEventEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*TriggerEventPayload)
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
		return s.TriggerEvent(ctx, p)
	}
}

// NewListEventsEndpoint returns an endpoint function that calls the method
// "list_events" of service "webhooks".
func NewListEventsEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ListEventsPayload)
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
		return s.ListEvents(ctx, p)
	}
}

// NewReplayEventEndpoint returns an endpoint function that calls the method
// "replay_event" of service "webhooks".
func NewReplayEventEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ReplayEventPayload)
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
		return s.ReplayEvent(ctx, p)
	}
}

// NewReceiveEndpoint returns an endpoint function that calls the method
// "receive" of service "webhooks".
func NewReceiveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*ReceivePayload)
		return s.Receive(ctx, p)
	}
}
