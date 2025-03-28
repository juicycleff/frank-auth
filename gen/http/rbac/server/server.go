// Code generated by goa v3.20.0, DO NOT EDIT.
//
// rbac HTTP server
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package server

import (
	"context"
	"net/http"
	"regexp"

	rbac "github.com/juicycleff/frank/gen/rbac"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the rbac service endpoint HTTP handlers.
type Server struct {
	Mounts               []*MountPoint
	ListPermissions      http.Handler
	CreatePermission     http.Handler
	GetPermission        http.Handler
	UpdatePermission     http.Handler
	DeletePermission     http.Handler
	ListRoles            http.Handler
	CreateRole           http.Handler
	GetRole              http.Handler
	UpdateRole           http.Handler
	DeleteRole           http.Handler
	ListRolePermissions  http.Handler
	AddRolePermission    http.Handler
	RemoveRolePermission http.Handler
	CheckPermission      http.Handler
	CheckRole            http.Handler
	CORS                 http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the rbac service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *rbac.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"ListPermissions", "GET", "/v1/permissions"},
			{"CreatePermission", "POST", "/v1/permissions"},
			{"GetPermission", "GET", "/v1/permissions/{id}"},
			{"UpdatePermission", "PUT", "/v1/permissions/{id}"},
			{"DeletePermission", "DELETE", "/v1/permissions/{id}"},
			{"ListRoles", "GET", "/v1/roles"},
			{"CreateRole", "POST", "/v1/roles"},
			{"GetRole", "GET", "/v1/roles/{id}"},
			{"UpdateRole", "PUT", "/v1/roles/{id}"},
			{"DeleteRole", "DELETE", "/v1/roles/{id}"},
			{"ListRolePermissions", "GET", "/v1/roles/{id}/permissions"},
			{"AddRolePermission", "POST", "/v1/roles/{id}/permissions"},
			{"RemoveRolePermission", "DELETE", "/v1/roles/{id}/permissions/{permission_id}"},
			{"CheckPermission", "GET", "/v1/access/check"},
			{"CheckRole", "GET", "/v1/access/check-role"},
			{"CORS", "OPTIONS", "/v1/permissions"},
			{"CORS", "OPTIONS", "/v1/permissions/{id}"},
			{"CORS", "OPTIONS", "/v1/roles"},
			{"CORS", "OPTIONS", "/v1/roles/{id}"},
			{"CORS", "OPTIONS", "/v1/roles/{id}/permissions"},
			{"CORS", "OPTIONS", "/v1/roles/{id}/permissions/{permission_id}"},
			{"CORS", "OPTIONS", "/v1/access/check"},
			{"CORS", "OPTIONS", "/v1/access/check-role"},
		},
		ListPermissions:      NewListPermissionsHandler(e.ListPermissions, mux, decoder, encoder, errhandler, formatter),
		CreatePermission:     NewCreatePermissionHandler(e.CreatePermission, mux, decoder, encoder, errhandler, formatter),
		GetPermission:        NewGetPermissionHandler(e.GetPermission, mux, decoder, encoder, errhandler, formatter),
		UpdatePermission:     NewUpdatePermissionHandler(e.UpdatePermission, mux, decoder, encoder, errhandler, formatter),
		DeletePermission:     NewDeletePermissionHandler(e.DeletePermission, mux, decoder, encoder, errhandler, formatter),
		ListRoles:            NewListRolesHandler(e.ListRoles, mux, decoder, encoder, errhandler, formatter),
		CreateRole:           NewCreateRoleHandler(e.CreateRole, mux, decoder, encoder, errhandler, formatter),
		GetRole:              NewGetRoleHandler(e.GetRole, mux, decoder, encoder, errhandler, formatter),
		UpdateRole:           NewUpdateRoleHandler(e.UpdateRole, mux, decoder, encoder, errhandler, formatter),
		DeleteRole:           NewDeleteRoleHandler(e.DeleteRole, mux, decoder, encoder, errhandler, formatter),
		ListRolePermissions:  NewListRolePermissionsHandler(e.ListRolePermissions, mux, decoder, encoder, errhandler, formatter),
		AddRolePermission:    NewAddRolePermissionHandler(e.AddRolePermission, mux, decoder, encoder, errhandler, formatter),
		RemoveRolePermission: NewRemoveRolePermissionHandler(e.RemoveRolePermission, mux, decoder, encoder, errhandler, formatter),
		CheckPermission:      NewCheckPermissionHandler(e.CheckPermission, mux, decoder, encoder, errhandler, formatter),
		CheckRole:            NewCheckRoleHandler(e.CheckRole, mux, decoder, encoder, errhandler, formatter),
		CORS:                 NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "rbac" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.ListPermissions = m(s.ListPermissions)
	s.CreatePermission = m(s.CreatePermission)
	s.GetPermission = m(s.GetPermission)
	s.UpdatePermission = m(s.UpdatePermission)
	s.DeletePermission = m(s.DeletePermission)
	s.ListRoles = m(s.ListRoles)
	s.CreateRole = m(s.CreateRole)
	s.GetRole = m(s.GetRole)
	s.UpdateRole = m(s.UpdateRole)
	s.DeleteRole = m(s.DeleteRole)
	s.ListRolePermissions = m(s.ListRolePermissions)
	s.AddRolePermission = m(s.AddRolePermission)
	s.RemoveRolePermission = m(s.RemoveRolePermission)
	s.CheckPermission = m(s.CheckPermission)
	s.CheckRole = m(s.CheckRole)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return rbac.MethodNames[:] }

// Mount configures the mux to serve the rbac endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountListPermissionsHandler(mux, h.ListPermissions)
	MountCreatePermissionHandler(mux, h.CreatePermission)
	MountGetPermissionHandler(mux, h.GetPermission)
	MountUpdatePermissionHandler(mux, h.UpdatePermission)
	MountDeletePermissionHandler(mux, h.DeletePermission)
	MountListRolesHandler(mux, h.ListRoles)
	MountCreateRoleHandler(mux, h.CreateRole)
	MountGetRoleHandler(mux, h.GetRole)
	MountUpdateRoleHandler(mux, h.UpdateRole)
	MountDeleteRoleHandler(mux, h.DeleteRole)
	MountListRolePermissionsHandler(mux, h.ListRolePermissions)
	MountAddRolePermissionHandler(mux, h.AddRolePermission)
	MountRemoveRolePermissionHandler(mux, h.RemoveRolePermission)
	MountCheckPermissionHandler(mux, h.CheckPermission)
	MountCheckRoleHandler(mux, h.CheckRole)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the rbac endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountListPermissionsHandler configures the mux to serve the "rbac" service
// "list_permissions" endpoint.
func MountListPermissionsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/permissions", f)
}

// NewListPermissionsHandler creates a HTTP handler which loads the HTTP
// request and calls the "rbac" service "list_permissions" endpoint.
func NewListPermissionsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListPermissionsRequest(mux, decoder)
		encodeResponse = EncodeListPermissionsResponse(encoder)
		encodeError    = EncodeListPermissionsError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list_permissions")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreatePermissionHandler configures the mux to serve the "rbac" service
// "create_permission" endpoint.
func MountCreatePermissionHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/permissions", f)
}

// NewCreatePermissionHandler creates a HTTP handler which loads the HTTP
// request and calls the "rbac" service "create_permission" endpoint.
func NewCreatePermissionHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreatePermissionRequest(mux, decoder)
		encodeResponse = EncodeCreatePermissionResponse(encoder)
		encodeError    = EncodeCreatePermissionError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_permission")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetPermissionHandler configures the mux to serve the "rbac" service
// "get_permission" endpoint.
func MountGetPermissionHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/permissions/{id}", f)
}

// NewGetPermissionHandler creates a HTTP handler which loads the HTTP request
// and calls the "rbac" service "get_permission" endpoint.
func NewGetPermissionHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetPermissionRequest(mux, decoder)
		encodeResponse = EncodeGetPermissionResponse(encoder)
		encodeError    = EncodeGetPermissionError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get_permission")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdatePermissionHandler configures the mux to serve the "rbac" service
// "update_permission" endpoint.
func MountUpdatePermissionHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/v1/permissions/{id}", f)
}

// NewUpdatePermissionHandler creates a HTTP handler which loads the HTTP
// request and calls the "rbac" service "update_permission" endpoint.
func NewUpdatePermissionHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdatePermissionRequest(mux, decoder)
		encodeResponse = EncodeUpdatePermissionResponse(encoder)
		encodeError    = EncodeUpdatePermissionError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update_permission")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeletePermissionHandler configures the mux to serve the "rbac" service
// "delete_permission" endpoint.
func MountDeletePermissionHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/v1/permissions/{id}", f)
}

// NewDeletePermissionHandler creates a HTTP handler which loads the HTTP
// request and calls the "rbac" service "delete_permission" endpoint.
func NewDeletePermissionHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeletePermissionRequest(mux, decoder)
		encodeResponse = EncodeDeletePermissionResponse(encoder)
		encodeError    = EncodeDeletePermissionError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete_permission")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListRolesHandler configures the mux to serve the "rbac" service
// "list_roles" endpoint.
func MountListRolesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/roles", f)
}

// NewListRolesHandler creates a HTTP handler which loads the HTTP request and
// calls the "rbac" service "list_roles" endpoint.
func NewListRolesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListRolesRequest(mux, decoder)
		encodeResponse = EncodeListRolesResponse(encoder)
		encodeError    = EncodeListRolesError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list_roles")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateRoleHandler configures the mux to serve the "rbac" service
// "create_role" endpoint.
func MountCreateRoleHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/roles", f)
}

// NewCreateRoleHandler creates a HTTP handler which loads the HTTP request and
// calls the "rbac" service "create_role" endpoint.
func NewCreateRoleHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateRoleRequest(mux, decoder)
		encodeResponse = EncodeCreateRoleResponse(encoder)
		encodeError    = EncodeCreateRoleError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create_role")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetRoleHandler configures the mux to serve the "rbac" service
// "get_role" endpoint.
func MountGetRoleHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/roles/{id}", f)
}

// NewGetRoleHandler creates a HTTP handler which loads the HTTP request and
// calls the "rbac" service "get_role" endpoint.
func NewGetRoleHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetRoleRequest(mux, decoder)
		encodeResponse = EncodeGetRoleResponse(encoder)
		encodeError    = EncodeGetRoleError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get_role")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateRoleHandler configures the mux to serve the "rbac" service
// "update_role" endpoint.
func MountUpdateRoleHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/v1/roles/{id}", f)
}

// NewUpdateRoleHandler creates a HTTP handler which loads the HTTP request and
// calls the "rbac" service "update_role" endpoint.
func NewUpdateRoleHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateRoleRequest(mux, decoder)
		encodeResponse = EncodeUpdateRoleResponse(encoder)
		encodeError    = EncodeUpdateRoleError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update_role")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteRoleHandler configures the mux to serve the "rbac" service
// "delete_role" endpoint.
func MountDeleteRoleHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/v1/roles/{id}", f)
}

// NewDeleteRoleHandler creates a HTTP handler which loads the HTTP request and
// calls the "rbac" service "delete_role" endpoint.
func NewDeleteRoleHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteRoleRequest(mux, decoder)
		encodeResponse = EncodeDeleteRoleResponse(encoder)
		encodeError    = EncodeDeleteRoleError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete_role")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountListRolePermissionsHandler configures the mux to serve the "rbac"
// service "list_role_permissions" endpoint.
func MountListRolePermissionsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/roles/{id}/permissions", f)
}

// NewListRolePermissionsHandler creates a HTTP handler which loads the HTTP
// request and calls the "rbac" service "list_role_permissions" endpoint.
func NewListRolePermissionsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListRolePermissionsRequest(mux, decoder)
		encodeResponse = EncodeListRolePermissionsResponse(encoder)
		encodeError    = EncodeListRolePermissionsError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list_role_permissions")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountAddRolePermissionHandler configures the mux to serve the "rbac" service
// "add_role_permission" endpoint.
func MountAddRolePermissionHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/roles/{id}/permissions", f)
}

// NewAddRolePermissionHandler creates a HTTP handler which loads the HTTP
// request and calls the "rbac" service "add_role_permission" endpoint.
func NewAddRolePermissionHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAddRolePermissionRequest(mux, decoder)
		encodeResponse = EncodeAddRolePermissionResponse(encoder)
		encodeError    = EncodeAddRolePermissionError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add_role_permission")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountRemoveRolePermissionHandler configures the mux to serve the "rbac"
// service "remove_role_permission" endpoint.
func MountRemoveRolePermissionHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/v1/roles/{id}/permissions/{permission_id}", f)
}

// NewRemoveRolePermissionHandler creates a HTTP handler which loads the HTTP
// request and calls the "rbac" service "remove_role_permission" endpoint.
func NewRemoveRolePermissionHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRemoveRolePermissionRequest(mux, decoder)
		encodeResponse = EncodeRemoveRolePermissionResponse(encoder)
		encodeError    = EncodeRemoveRolePermissionError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "remove_role_permission")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCheckPermissionHandler configures the mux to serve the "rbac" service
// "check_permission" endpoint.
func MountCheckPermissionHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/access/check", f)
}

// NewCheckPermissionHandler creates a HTTP handler which loads the HTTP
// request and calls the "rbac" service "check_permission" endpoint.
func NewCheckPermissionHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCheckPermissionRequest(mux, decoder)
		encodeResponse = EncodeCheckPermissionResponse(encoder)
		encodeError    = EncodeCheckPermissionError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "check_permission")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCheckRoleHandler configures the mux to serve the "rbac" service
// "check_role" endpoint.
func MountCheckRoleHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleRbacOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/access/check-role", f)
}

// NewCheckRoleHandler creates a HTTP handler which loads the HTTP request and
// calls the "rbac" service "check_role" endpoint.
func NewCheckRoleHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCheckRoleRequest(mux, decoder)
		encodeResponse = EncodeCheckRoleResponse(encoder)
		encodeError    = EncodeCheckRoleError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "check_role")
		ctx = context.WithValue(ctx, goa.ServiceKey, "rbac")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service rbac.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleRbacOrigin(h)
	mux.Handle("OPTIONS", "/v1/permissions", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/permissions/{id}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/roles", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/roles/{id}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/roles/{id}/permissions", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/roles/{id}/permissions/{permission_id}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/access/check", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/access/check-role", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 204 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
}

// HandleRbacOrigin applies the CORS response headers corresponding to the
// origin for the service rbac.
func HandleRbacOrigin(h http.Handler) http.Handler {
	spec1 := regexp.MustCompile(".*localhost.*")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "*.frank.com") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", "100")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Headers", "X-Shared-Secret, X-Api-Version")
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec1) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "X-Request-Id")
			w.Header().Set("Access-Control-Max-Age", "600")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOrigin(origin, "localhost") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.WriteHeader(204)
				return
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
