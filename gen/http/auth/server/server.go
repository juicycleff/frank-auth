// Code generated by goa v3.20.0, DO NOT EDIT.
//
// auth HTTP server
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package server

import (
	"context"
	"net/http"
	"regexp"

	auth "github.com/juicycleff/frank/gen/auth"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the auth service endpoint HTTP handlers.
type Server struct {
	Mounts         []*MountPoint
	Login          http.Handler
	Register       http.Handler
	Logout         http.Handler
	RefreshToken   http.Handler
	ForgotPassword http.Handler
	ResetPassword  http.Handler
	VerifyEmail    http.Handler
	Me             http.Handler
	Csrf           http.Handler
	CORS           http.Handler
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

// New instantiates HTTP handlers for all the auth service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *auth.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Login", "POST", "/v1/auth/login"},
			{"Register", "POST", "/v1/auth/register"},
			{"Logout", "POST", "/v1/auth/logout"},
			{"RefreshToken", "POST", "/v1/auth/refresh"},
			{"ForgotPassword", "POST", "/v1/auth/forgot-password"},
			{"ResetPassword", "POST", "/v1/auth/reset-password"},
			{"VerifyEmail", "POST", "/v1/auth/verify-email"},
			{"Me", "GET", "/v1/auth/me"},
			{"Csrf", "GET", "/v1/auth/csrf-token"},
			{"CORS", "OPTIONS", "/v1/auth/login"},
			{"CORS", "OPTIONS", "/v1/auth/register"},
			{"CORS", "OPTIONS", "/v1/auth/logout"},
			{"CORS", "OPTIONS", "/v1/auth/refresh"},
			{"CORS", "OPTIONS", "/v1/auth/forgot-password"},
			{"CORS", "OPTIONS", "/v1/auth/reset-password"},
			{"CORS", "OPTIONS", "/v1/auth/verify-email"},
			{"CORS", "OPTIONS", "/v1/auth/me"},
			{"CORS", "OPTIONS", "/v1/auth/csrf-token"},
		},
		Login:          NewLoginHandler(e.Login, mux, decoder, encoder, errhandler, formatter),
		Register:       NewRegisterHandler(e.Register, mux, decoder, encoder, errhandler, formatter),
		Logout:         NewLogoutHandler(e.Logout, mux, decoder, encoder, errhandler, formatter),
		RefreshToken:   NewRefreshTokenHandler(e.RefreshToken, mux, decoder, encoder, errhandler, formatter),
		ForgotPassword: NewForgotPasswordHandler(e.ForgotPassword, mux, decoder, encoder, errhandler, formatter),
		ResetPassword:  NewResetPasswordHandler(e.ResetPassword, mux, decoder, encoder, errhandler, formatter),
		VerifyEmail:    NewVerifyEmailHandler(e.VerifyEmail, mux, decoder, encoder, errhandler, formatter),
		Me:             NewMeHandler(e.Me, mux, decoder, encoder, errhandler, formatter),
		Csrf:           NewCsrfHandler(e.Csrf, mux, decoder, encoder, errhandler, formatter),
		CORS:           NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "auth" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Login = m(s.Login)
	s.Register = m(s.Register)
	s.Logout = m(s.Logout)
	s.RefreshToken = m(s.RefreshToken)
	s.ForgotPassword = m(s.ForgotPassword)
	s.ResetPassword = m(s.ResetPassword)
	s.VerifyEmail = m(s.VerifyEmail)
	s.Me = m(s.Me)
	s.Csrf = m(s.Csrf)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return auth.MethodNames[:] }

// Mount configures the mux to serve the auth endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountLoginHandler(mux, h.Login)
	MountRegisterHandler(mux, h.Register)
	MountLogoutHandler(mux, h.Logout)
	MountRefreshTokenHandler(mux, h.RefreshToken)
	MountForgotPasswordHandler(mux, h.ForgotPassword)
	MountResetPasswordHandler(mux, h.ResetPassword)
	MountVerifyEmailHandler(mux, h.VerifyEmail)
	MountMeHandler(mux, h.Me)
	MountCsrfHandler(mux, h.Csrf)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the auth endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountLoginHandler configures the mux to serve the "auth" service "login"
// endpoint.
func MountLoginHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/login", f)
}

// NewLoginHandler creates a HTTP handler which loads the HTTP request and
// calls the "auth" service "login" endpoint.
func NewLoginHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeLoginRequest(mux, decoder)
		encodeResponse = EncodeLoginResponse(encoder)
		encodeError    = EncodeLoginError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "login")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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

// MountRegisterHandler configures the mux to serve the "auth" service
// "register" endpoint.
func MountRegisterHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/register", f)
}

// NewRegisterHandler creates a HTTP handler which loads the HTTP request and
// calls the "auth" service "register" endpoint.
func NewRegisterHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRegisterRequest(mux, decoder)
		encodeResponse = EncodeRegisterResponse(encoder)
		encodeError    = EncodeRegisterError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "register")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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

// MountLogoutHandler configures the mux to serve the "auth" service "logout"
// endpoint.
func MountLogoutHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/logout", f)
}

// NewLogoutHandler creates a HTTP handler which loads the HTTP request and
// calls the "auth" service "logout" endpoint.
func NewLogoutHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeLogoutRequest(mux, decoder)
		encodeResponse = EncodeLogoutResponse(encoder)
		encodeError    = EncodeLogoutError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "logout")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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

// MountRefreshTokenHandler configures the mux to serve the "auth" service
// "refresh_token" endpoint.
func MountRefreshTokenHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/refresh", f)
}

// NewRefreshTokenHandler creates a HTTP handler which loads the HTTP request
// and calls the "auth" service "refresh_token" endpoint.
func NewRefreshTokenHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRefreshTokenRequest(mux, decoder)
		encodeResponse = EncodeRefreshTokenResponse(encoder)
		encodeError    = EncodeRefreshTokenError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "refresh_token")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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

// MountForgotPasswordHandler configures the mux to serve the "auth" service
// "forgot_password" endpoint.
func MountForgotPasswordHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/forgot-password", f)
}

// NewForgotPasswordHandler creates a HTTP handler which loads the HTTP request
// and calls the "auth" service "forgot_password" endpoint.
func NewForgotPasswordHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeForgotPasswordRequest(mux, decoder)
		encodeResponse = EncodeForgotPasswordResponse(encoder)
		encodeError    = EncodeForgotPasswordError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "forgot_password")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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

// MountResetPasswordHandler configures the mux to serve the "auth" service
// "reset_password" endpoint.
func MountResetPasswordHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/reset-password", f)
}

// NewResetPasswordHandler creates a HTTP handler which loads the HTTP request
// and calls the "auth" service "reset_password" endpoint.
func NewResetPasswordHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeResetPasswordRequest(mux, decoder)
		encodeResponse = EncodeResetPasswordResponse(encoder)
		encodeError    = EncodeResetPasswordError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "reset_password")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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

// MountVerifyEmailHandler configures the mux to serve the "auth" service
// "verify_email" endpoint.
func MountVerifyEmailHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/v1/auth/verify-email", f)
}

// NewVerifyEmailHandler creates a HTTP handler which loads the HTTP request
// and calls the "auth" service "verify_email" endpoint.
func NewVerifyEmailHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeVerifyEmailRequest(mux, decoder)
		encodeResponse = EncodeVerifyEmailResponse(encoder)
		encodeError    = EncodeVerifyEmailError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "verify_email")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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

// MountMeHandler configures the mux to serve the "auth" service "me" endpoint.
func MountMeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/auth/me", f)
}

// NewMeHandler creates a HTTP handler which loads the HTTP request and calls
// the "auth" service "me" endpoint.
func NewMeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeMeRequest(mux, decoder)
		encodeResponse = EncodeMeResponse(encoder)
		encodeError    = EncodeMeError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "me")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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

// MountCsrfHandler configures the mux to serve the "auth" service "csrf"
// endpoint.
func MountCsrfHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAuthOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/auth/csrf-token", f)
}

// NewCsrfHandler creates a HTTP handler which loads the HTTP request and calls
// the "auth" service "csrf" endpoint.
func NewCsrfHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCsrfRequest(mux, decoder)
		encodeResponse = EncodeCsrfResponse(encoder)
		encodeError    = EncodeCsrfError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "csrf")
		ctx = context.WithValue(ctx, goa.ServiceKey, "auth")
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
// service auth.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleAuthOrigin(h)
	mux.Handle("OPTIONS", "/v1/auth/login", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/register", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/logout", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/refresh", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/forgot-password", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/reset-password", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/verify-email", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/me", h.ServeHTTP)
	mux.Handle("OPTIONS", "/v1/auth/csrf-token", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 204 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
}

// HandleAuthOrigin applies the CORS response headers corresponding to the
// origin for the service auth.
func HandleAuthOrigin(h http.Handler) http.Handler {
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
