// Code generated by goa v3.20.0, DO NOT EDIT.
//
// admin HTTP server
//
// Command:
// $ goa gen github.com/juicycleff/frank/design -o .

package server

import (
	"context"
	"net/http"
	"path"
	"regexp"

	admin "github.com/juicycleff/frank/gen/admin"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the admin service endpoint HTTP handlers.
type Server struct {
	Mounts                []*MountPoint
	Home                  http.Handler
	CORS                  http.Handler
	UIBuild               http.Handler
	AdminPublicRobotsTxt  http.Handler
	AdminPublicFaviconIco http.Handler
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

// New instantiates HTTP handlers for all the admin service endpoints using the
// provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *admin.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemUIBuild http.FileSystem,
	fileSystemAdminPublicRobotsTxt http.FileSystem,
	fileSystemAdminPublicFaviconIco http.FileSystem,
) *Server {
	if fileSystemUIBuild == nil {
		fileSystemUIBuild = http.Dir(".")
	}
	fileSystemUIBuild = appendPrefix(fileSystemUIBuild, "/ui/build")
	if fileSystemAdminPublicRobotsTxt == nil {
		fileSystemAdminPublicRobotsTxt = http.Dir(".")
	}
	fileSystemAdminPublicRobotsTxt = appendPrefix(fileSystemAdminPublicRobotsTxt, "/admin/public")
	if fileSystemAdminPublicFaviconIco == nil {
		fileSystemAdminPublicFaviconIco = http.Dir(".")
	}
	fileSystemAdminPublicFaviconIco = appendPrefix(fileSystemAdminPublicFaviconIco, "/admin/public")
	return &Server{
		Mounts: []*MountPoint{
			{"Home", "GET", "/"},
			{"Home", "GET", "/admin"},
			{"CORS", "OPTIONS", "/"},
			{"CORS", "OPTIONS", "/admin"},
			{"CORS", "OPTIONS", "/admin/{*filepath}"},
			{"CORS", "OPTIONS", "/robots.txt"},
			{"CORS", "OPTIONS", "/favicon.ico"},
			{"Serve ui/build", "GET", "/admin"},
			{"Serve admin/public/robots.txt", "GET", "/robots.txt"},
			{"Serve admin/public/favicon.ico", "GET", "/favicon.ico"},
		},
		Home:                  NewHomeHandler(e.Home, mux, decoder, encoder, errhandler, formatter),
		CORS:                  NewCORSHandler(),
		UIBuild:               http.FileServer(fileSystemUIBuild),
		AdminPublicRobotsTxt:  http.FileServer(fileSystemAdminPublicRobotsTxt),
		AdminPublicFaviconIco: http.FileServer(fileSystemAdminPublicFaviconIco),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "admin" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Home = m(s.Home)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return admin.MethodNames[:] }

// Mount configures the mux to serve the admin endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountHomeHandler(mux, h.Home)
	MountCORSHandler(mux, h.CORS)
	MountUIBuild(mux, http.StripPrefix("/admin", h.UIBuild))
	MountAdminPublicRobotsTxt(mux, h.AdminPublicRobotsTxt)
	MountAdminPublicFaviconIco(mux, h.AdminPublicFaviconIco)
}

// Mount configures the mux to serve the admin endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountHomeHandler configures the mux to serve the "admin" service "home"
// endpoint.
func MountHomeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandleAdminOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/", f)
	mux.Handle("GET", "/admin", f)
}

// NewHomeHandler creates a HTTP handler which loads the HTTP request and calls
// the "admin" service "home" endpoint.
func NewHomeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "home")
		ctx = context.WithValue(ctx, goa.ServiceKey, "admin")
		http.Redirect(w, r, "/admin/", http.StatusMovedPermanently)
	})
}

// appendFS is a custom implementation of fs.FS that appends a specified prefix
// to the file paths before delegating the Open call to the underlying fs.FS.
type appendFS struct {
	prefix string
	fs     http.FileSystem
}

// Open opens the named file, appending the prefix to the file path before
// passing it to the underlying fs.FS.
func (s appendFS) Open(name string) (http.File, error) {
	switch name {
	}
	return s.fs.Open(path.Join(s.prefix, name))
}

// appendPrefix returns a new fs.FS that appends the specified prefix to file paths
// before delegating to the provided embed.FS.
func appendPrefix(fsys http.FileSystem, prefix string) http.FileSystem {
	return appendFS{prefix: prefix, fs: fsys}
}

// MountUIBuild configures the mux to serve GET request made to "/admin".
func MountUIBuild(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/admin/", HandleAdminOrigin(h).ServeHTTP)
	mux.Handle("GET", "/admin/{*filepath}", HandleAdminOrigin(h).ServeHTTP)
}

// MountAdminPublicRobotsTxt configures the mux to serve GET request made to
// "/robots.txt".
func MountAdminPublicRobotsTxt(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/robots.txt", HandleAdminOrigin(h).ServeHTTP)
}

// MountAdminPublicFaviconIco configures the mux to serve GET request made to
// "/favicon.ico".
func MountAdminPublicFaviconIco(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/favicon.ico", HandleAdminOrigin(h).ServeHTTP)
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service admin.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandleAdminOrigin(h)
	mux.Handle("OPTIONS", "/", h.ServeHTTP)
	mux.Handle("OPTIONS", "/admin", h.ServeHTTP)
	mux.Handle("OPTIONS", "/admin/{*filepath}", h.ServeHTTP)
	mux.Handle("OPTIONS", "/robots.txt", h.ServeHTTP)
	mux.Handle("OPTIONS", "/favicon.ico", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 204 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
	})
}

// HandleAdminOrigin applies the CORS response headers corresponding to the
// origin for the service admin.
func HandleAdminOrigin(h http.Handler) http.Handler {
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
