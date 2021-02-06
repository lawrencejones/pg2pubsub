// Code generated by goa v3.2.6, DO NOT EDIT.
//
// Health HTTP server
//
// Command:
// $ goa gen github.com/lawrencejones/pgsink/api/design -o api

package server

import (
	"context"
	"net/http"

	health "github.com/lawrencejones/pgsink/api/gen/health"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the Health service endpoint HTTP handlers.
type Server struct {
	Mounts []*MountPoint
	Check  http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
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

// New instantiates HTTP handlers for all the Health service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *health.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Check", "GET", "/api/health/check"},
		},
		Check: NewCheckHandler(e.Check, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "Health" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Check = m(s.Check)
}

// Mount configures the mux to serve the Health endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCheckHandler(mux, h.Check)
}

// MountCheckHandler configures the mux to serve the "Health" service "Check"
// endpoint.
func MountCheckHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/health/check", f)
}

// NewCheckHandler creates a HTTP handler which loads the HTTP request and
// calls the "Health" service "Check" endpoint.
func NewCheckHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		encodeResponse = EncodeCheckResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Check")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Health")
		var err error
		res, err := endpoint(ctx, nil)
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
