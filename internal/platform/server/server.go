package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marcohb99/go-api-example/internal/creating"
	"github.com/marcohb99/go-api-example/internal/platform/server/handler/health"
	"github.com/marcohb99/go-api-example/internal/platform/server/handler/hello"
	"github.com/marcohb99/go-api-example/internal/platform/server/handler/release"
)

// Server encapsulates a server with an engine and an address
type Server struct {
	httpAddr string
	engine   *gin.Engine

	// dependencies
	creatingReleaseServcie creating.ReleaseService
}

func New(host string, port uint, creatingReleaseServcie creating.ReleaseService) Server {
	srv := Server{
		httpAddr: fmt.Sprintf("%s:%d", host, port),
		engine:   gin.New(),

		// dependencies
		creatingReleaseServcie: creatingReleaseServcie,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

// ROUTES
func (s *Server) registerRoutes() {
	// hc
	s.engine.GET("/health", health.CheckHandler())

	// hello
	s.engine.GET("/hello", hello.GetHandler())

	// release
	s.engine.POST("/releases", release.CreateHandler(s.creatingReleaseServcie))
}
