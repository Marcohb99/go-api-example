package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
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
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
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
