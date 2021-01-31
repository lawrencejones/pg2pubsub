package cmd

import (
	"net/http"
	"os"

	kitlog "github.com/go-kit/kit/log"
	"github.com/lawrencejones/pgsink/api/gen/health"
	healthserver "github.com/lawrencejones/pgsink/api/gen/http/health/server"
	importsserver "github.com/lawrencejones/pgsink/api/gen/http/imports/server"
	tablesserver "github.com/lawrencejones/pgsink/api/gen/http/tables/server"
	"github.com/lawrencejones/pgsink/api/gen/imports"
	"github.com/lawrencejones/pgsink/api/gen/tables"
	middleware "github.com/lawrencejones/pgsink/internal/middleware"
	goahttp "goa.design/goa/v3/http"
	httpmw "goa.design/goa/v3/http/middleware"
)

// buildHTTPServer instantiates a new HTTP server with Goa endpoints
func buildHTTPServer(logger kitlog.Logger, addr string, tablesEndpoints *tables.Endpoints, importsEndpoints *imports.Endpoints, healthEndpoints *health.Endpoints, debug bool) *http.Server {
	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer = goahttp.NewMuxer()

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		tablesServer  *tablesserver.Server
		importsServer *importsserver.Server
		healthServer  *healthserver.Server
	)
	{
		tablesServer = tablesserver.New(tablesEndpoints, mux, dec, enc, nil, nil)
		importsServer = importsserver.New(importsEndpoints, mux, dec, enc, nil, nil)
		healthServer = healthserver.New(healthEndpoints, mux, dec, enc, nil, nil)
		if debug {
			servers := goahttp.Servers{
				tablesServer,
				importsServer,
				healthServer,
			}
			servers.Use(httpmw.Debug(mux, os.Stderr))
		}

		// Mount the server
		tablesserver.Mount(mux, tablesServer)
		importsserver.Mount(mux, importsServer)
		healthserver.Mount(mux, healthServer)
	}

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints, and run in reverse order.
	var handler http.Handler = mux
	{
		// Default observability, request logging etc
		handler = middleware.ObserveHTTP(logger)(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: addr, Handler: handler}
	for _, m := range tablesServer.Mounts {
		logger.Log("event", "mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}
	for _, m := range importsServer.Mounts {
		logger.Log("event", "mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}
	for _, m := range healthServer.Mounts {
		logger.Log("event", "mount", "method", m.Method, "verb", m.Verb, "pattern", m.Pattern)
	}

	return srv
}