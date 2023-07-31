package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/nownabe/golink/api/gen/golink/v1/golinkv1connect"
)

const (
	readHeaderTimeoutSeconds = 10
	shutdownTimeoutSeconds   = 120
)

type API interface {
	Run() error
}

func New(
	golinkSvc golinkv1connect.GolinkServiceHandler,
	port, pathPrefix string,
	allowedOrigins []string,
) API {
	return &api{
		golinkSvc:      golinkSvc,
		port:           port,
		pathPrefix:     pathPrefix,
		allowedOrigins: allowedOrigins,
	}
}

type api struct {
	golinkSvc golinkv1connect.GolinkServiceHandler

	port           string
	pathPrefix     string
	allowedOrigins []string
}

func (a *api) Run() error {
	return a.serve()
}

func (a *api) serve() error {
	path, handler := golinkv1connect.NewGolinkServiceHandler(a.golinkSvc)

	mux := http.NewServeMux()
	mux.Handle(a.pathPrefix+path, a.trimPrefix(handler))
	mux.HandleFunc(a.pathPrefix+"/healthz", a.healthz)
	mux.HandleFunc("/", http.NotFound)

	h2s := &http2.Server{}
	h1s := &http.Server{
		Addr:              ":" + a.port,
		Handler:           a.cors(h2c.NewHandler(mux, h2s)),
		ReadHeaderTimeout: readHeaderTimeoutSeconds * time.Second,
	}

	idleConnsClosed := make(chan struct{})

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

		sig := <-ch
		log.Print(sig)

		// TODO: log: info: received signal and terminating

		ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeoutSeconds*time.Second)
		defer cancel()

		if err := h1s.Shutdown(ctx); err != nil {
			// TODO: log: err: failed to shutdown gracefully
			log.Print(err)
		}

		// TODO: log: info: completed shutdown gracefully

		close(idleConnsClosed)
	}()

	// TODO: log: info: started

	if err := h1s.ListenAndServe(); err != http.ErrServerClosed {
		// TODO: log: fatal: failed to listen and serve
		return err
	}

	<-idleConnsClosed

	// TODO: log: info: bye
	return nil
}

func (a *api) healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// connect-go doesn't support path prefix.
// So we need to trim it.
// See: https://github.com/bufbuild/connect-go/blob/843d045a5a76ee6236ecd5f05320f58446afec26/cmd/protoc-gen-connect-go/main.go#L215
func (a *api) trimPrefix(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		originalPath := r.URL.Path
		r.URL.Path = strings.TrimPrefix(r.URL.Path, a.pathPrefix)
		h.ServeHTTP(w, r)
		r.URL.Path = originalPath
	})
}

func (a *api) cors(h http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   a.allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	})

	return c.Handler(h)
}