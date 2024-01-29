package bootstrap

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/marcohb99/go-api-example/internal/creating"
	"github.com/marcohb99/go-api-example/internal/platform/bus/inmemory"
	"github.com/marcohb99/go-api-example/internal/platform/server"
	"github.com/marcohb99/go-api-example/internal/platform/storage/mysql"
)

const (
	host = "localhost"
	port = 8080

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

	srv := server.New(host, port, commandBus)
	return srv.Run()
}
