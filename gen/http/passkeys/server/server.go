// Code generated by goa v3.20.0, DO NOT EDIT.
//
// passkeys HTTP server
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package server

import (
	"context"
	"net/http"
	"regexp"

	passkeys "github.com/juicycleff/frank/gen/passkeys"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the passkeys service endpoint HTTP handlers.
type Server struct {
	Mounts           []*MountPoint
	RegisterBegin    http.Handler
	RegisterComplete http.Handler
	LoginBegin       http.Handler
	LoginComplete    http.Handler
	List             http.Handler
	Update           http.Handler
	Delete           http.Handler
	CORS             http.Handler
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

// New instantiates HTTP handlers for all the passkeys service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *passkeys.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"RegisterBegin", "POST", "/v1/auth/passkeys/register/begin"},
			{"RegisterComplete", "POST", "/v1/auth/passkeys/register/complete"},
			{"LoginBegin", "POST", "/v1/auth/passkeys/login/begin"},
			{"LoginComplete", "POST", "/v1/auth/passkeys/login/complete"},
			{"List", "GET", "/v1/auth/passkeys"},
			{"Update", "PUT", "/v1/auth/passkeys/{id}"},
			{"Delete", "DELETE", "/v1/auth/passkeys/{id}"},
			{"CORS", "OPTIONS", "/v1/auth/passkeys/register/begin"},
			{"CORS", "OPTIONS", "/v1/auth/passkeys/register/complete"},
			{"CORS", "OPTIONS", "/v1/auth/passkeys/login/begin"},
			{"CORS", "OPTIONS", "/v1/auth/passkeys/login/complete"},
			{"CORS", "OPTIONS", "/v1/auth/passkeys"},
			{"CORS", "OPTIONS", "/v1/auth/passkeys/{id}"},
		},
		RegisterBegin:    NewRegisterBeginHandler(e.RegisterBegin, mux, decoder, encoder, errhandler, formatter),
		RegisterComplete: NewRegisterCompleteHandler(e.RegisterComplete, mux, decoder, encoder, errhandler, formatter),
		LoginBegin:       NewLoginBeginHandler(e.LoginBegin, mux, decoder, encoder, errhandler, formatter),
		LoginComplete:    NewLoginCompleteHandler(e.LoginComplete, mux, decoder, encoder, errhandler, formatter),
		List:             NewListHandler(e.List, mux, decoder, encoder, errhandler, formatter),
		Update:           NewUpdateHandler(e.Update, mux, decoder, encoder, errhandler, formatter),
		Delete:           NewDeleteHandler(e.Delete, mux, decoder, encoder, errhandler, formatter),
		CORS:             NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "passkeys" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.RegisterBegin = m(s.RegisterBegin)
	s.RegisterComplete = m(s.RegisterComplete)
	s.LoginBegin = m(s.LoginBegin)
	s.LoginComplete = m(s.LoginComplete)
	s.List = m(s.List)
	s.Update = m(s.Update)
	s.Delete = m(s.Delete)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return passkeys.MethodNames[:] }

// Mount configures the mux to serve the passkeys endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountRegisterBeginHandler(mux, h.RegisterBegin)
	MountRegisterCompleteHandler(mux, h.RegisterComplete)
	MountLoginBeginHandler(mux, h.LoginBegin)
	MountLoginCompleteHandler(mux, h.LoginComplete)
	MountListHandler(mux, h.List)
	MountUpdateHandler(mux, h.Update)
	MountDeleteHandler(mux, h.Delete)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the passkeys endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountRegisterBeginHandler configures the mux to serve the "passkeys" service
// "register_begin" endpoint.
func MountRegisterBeginHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePasskeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/passkeys/register/begin", f)
}

// NewRegisterBeginHandler creates a HTTP handler which loads the HTTP request
// and calls the "passkeys" service "register_begin" endpoint.
func NewRegisterBeginHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRegisterBeginRequest(mux, decoder)
		encodeResponse = EncodeRegisterBeginResponse(encoder)
		encodeError    = EncodeRegisterBeginError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "register_begin")
		ctx = context.WithValue(ctx, goa.ServiceKey, "passkeys")
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

// MountRegisterCompleteHandler configures the mux to serve the "passkeys"
// service "register_complete" endpoint.
func MountRegisterCompleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePasskeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/passkeys/register/complete", f)
}

// NewRegisterCompleteHandler creates a HTTP handler which loads the HTTP
// request and calls the "passkeys" service "register_complete" endpoint.
func NewRegisterCompleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRegisterCompleteRequest(mux, decoder)
		encodeResponse = EncodeRegisterCompleteResponse(encoder)
		encodeError    = EncodeRegisterCompleteError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "register_complete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "passkeys")
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

// MountLoginBeginHandler configures the mux to serve the "passkeys" service
// "login_begin" endpoint.
func MountLoginBeginHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePasskeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/passkeys/login/begin", f)
}

// NewLoginBeginHandler creates a HTTP handler which loads the HTTP request and
// calls the "passkeys" service "login_begin" endpoint.
func NewLoginBeginHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeLoginBeginRequest(mux, decoder)
		encodeResponse = EncodeLoginBeginResponse(encoder)
		encodeError    = EncodeLoginBeginError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "login_begin")
		ctx = context.WithValue(ctx, goa.ServiceKey, "passkeys")
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

// MountLoginCompleteHandler configures the mux to serve the "passkeys" service
// "login_complete" endpoint.
func MountLoginCompleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePasskeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/passkeys/login/complete", f)
}

// NewLoginCompleteHandler creates a HTTP handler which loads the HTTP request
// and calls the "passkeys" service "login_complete" endpoint.
func NewLoginCompleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeLoginCompleteRequest(mux, decoder)
		encodeResponse = EncodeLoginCompleteResponse(encoder)
		encodeError    = EncodeLoginCompleteError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "login_complete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "passkeys")
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

// MountListHandler configures the mux to serve the "passkeys" service "list"
// endpoint.
func MountListHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePasskeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/auth/passkeys", f)
}

// NewListHandler creates a HTTP handler which loads the HTTP request and calls
// the "passkeys" service "list" endpoint.
func NewListHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeListRequest(mux, decoder)
		encodeResponse = EncodeListResponse(encoder)
		encodeError    = EncodeListError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "list")
		ctx = context.WithValue(ctx, goa.ServiceKey, "passkeys")
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

// MountUpdateHandler configures the mux to serve the "passkeys" service
// "update" endpoint.
func MountUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePasskeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/v1/auth/passkeys/{id}", f)
}

// NewUpdateHandler creates a HTTP handler which loads the HTTP request and
// calls the "passkeys" service "update" endpoint.
func NewUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateRequest(mux, decoder)
		encodeResponse = EncodeUpdateResponse(encoder)
		encodeError    = EncodeUpdateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "passkeys")
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

// MountDeleteHandler configures the mux to serve the "passkeys" service
// "delete" endpoint.
func MountDeleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePasskeysOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/v1/auth/passkeys/{id}", f)
}

// NewDeleteHandler creates a HTTP handler which loads the HTTP request and
// calls the "passkeys" service "delete" endpoint.
func NewDeleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteRequest(mux, decoder)
		encodeResponse = EncodeDeleteResponse(encoder)
		encodeError    = EncodeDeleteError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "passkeys")
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
// service passkeys.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandlePasskeysOrigin(h)
	mux.Handle("OPTIONS", "/v1/auth/passkeys/register/begin", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/passkeys/register/complete", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/passkeys/login/begin", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/passkeys/login/complete", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/passkeys", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/passkeys/{id}", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 204 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
}

// HandlePasskeysOrigin applies the CORS response headers corresponding to the
// origin for the service passkeys.
func HandlePasskeysOrigin(h http.Handler) http.Handler {
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
