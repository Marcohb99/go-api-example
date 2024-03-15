package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	apiExample "github.com/marcohb99/go-api-example/internal"
	"github.com/marcohb99/go-api-example/internal/increasing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcohb99/go-api-example/internal/creating"
	"github.com/marcohb99/go-api-example/internal/platform/bus/inmemory"
	"github.com/marcohb99/go-api-example/internal/platform/server"
	"github.com/marcohb99/go-api-example/internal/platform/storage/mysql"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

	// Database constants
	dbUser    = "mhb"
	dbPass    = "mhb"
	dbHost    = "localhost"
	dbPort    = "3306"
	dbName    = "releases"
	dbTimeout = 5 * time.Second
)

func Run() error {
	// MYSQL
	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	// BUSES
	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	// REPOSITORIES
	releaseRepository := mysql.NewReleaseRepository(db, dbTimeout)

	// SERVICES
	creatingReleaseService := creating.NewReleaseService(releaseRepository, eventBus)
	increasingReleaseCounterService := increasing.NewReleaseCounterService()

	// COMMANDS
	createReleaseCommandHandler := creating.NewReleaseCommandHandler(creatingReleaseService)
	commandBus.Register(creating.ReleaseCommandType, createReleaseCommandHandler)

	// EVENTS
	eventBus.Subscribe(
		apiExample.ReleaseCreatedEventType,
		creating.NewIncreaseReleasesCounterOnReleaseCreated(increasingReleaseCounterService),
	)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `required:"true" default:"localhost"`
	Port            uint          `required:"true" default:"8080"`
	ShutdownTimeout time.Duration `split_words:"true" default:"10s"`
	// Database configuration
	DbUser    string        `required:"true" default:"mhb"`
	DbPass    string        `required:"true" default:"mhb"`
	DbHost    string        `required:"true" default:"localhost"`
	DbPort    uint          `required:"true" default:"3306"`
	DbName    string        `required:"true" default:"mhb"`
	DbTimeout time.Duration `default:"5s"`
}
