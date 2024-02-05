package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcohb99/go-api-example/internal/creating"
	"github.com/marcohb99/go-api-example/internal/platform/bus/inmemory"
	"github.com/marcohb99/go-api-example/internal/platform/server"
	"github.com/marcohb99/go-api-example/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080
	shutdownTimeout = 10 * time.Second

	// Database constants
	dbUser = "mhb"
	dbPass = "mhb"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "releases"
)

func Run() error {
	// MySQL connection
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	// repository
	releaseRepository := mysql.NewReleaseRepository(db)

	// service
	creatingReleaseService := creating.NewReleaseService(releaseRepository)

	// command
	var (
		commandBus = inmemory.NewCommandBus()
	)
	createReleaseCommandHandler := creating.NewReleaseCommandHandler(creatingReleaseService)
	commandBus.Register(creating.ReleaseCommandType, createReleaseCommandHandler)

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}
