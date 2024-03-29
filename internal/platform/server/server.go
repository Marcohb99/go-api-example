package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marcohb99/go-api-example/internal/platform/server/handler/health"
	"github.com/marcohb99/go-api-example/internal/platform/server/handler/hello"
	"github.com/marcohb99/go-api-example/internal/platform/server/handler/release"
	"github.com/marcohb99/go-api-example/internal/platform/server/middleware/logging"
	"github.com/marcohb99/go-api-example/internal/platform/server/middleware/recovery"
	command "github.com/marcohb99/go-api-example/kit"
)

// Server encapsulates a server with an engine and an address
type Server struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	// dependencies
	commandBus command.Bus
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus) (context.Context, Server) {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),

		shutdownTimeout: shutdownTimeout,

		// dependencies
		commandBus: commandBus,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func serverContext(ctx context.Context) context.Context {
	// create a channel to listen for OS signals
	c := make(chan os.Signal, 1)
	// make the channel listen for an interrupt signal
	signal.Notify(c, os.Interrupt)
	// WithCancel returns a copy of parent with a new Done channel.
	// The returned context's Done channel is closed when the returned cancel function is called or
	// when the parent context's Done channel is closed, whichever happens first.
	// Canceling this context releases resources associated with it, so code should call cancel
	// as soon as the operations running in this Context complete.
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c // bloquea el canal hasta que se notifique una señal
		cancel()
	}()

	return ctx
}

func (s *Server) Run(ctx context.Context) error {
	log.Println("Server running on", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	// run the server in a goroutine so that it doesn't block
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done() // receive a signal to stop the server
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

// ROUTES
func (s *Server) registerRoutes() {
	s.engine.Use(recovery.Middleware(), logging.Middleware())
	// hc
	s.engine.GET("/health", health.CheckHandler())

	// hello
	s.engine.GET("/hello", hello.GetHandler())

	// release
	s.engine.GET("/releases", release.GetAllHandler(s.commandBus))
	s.engine.POST("/releases", release.CreateHandler(s.commandBus))
}
