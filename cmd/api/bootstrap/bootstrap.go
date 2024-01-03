package bootstrap

import (
	"database/sql"
	"fmt"

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

	srv := server.New(host, port, releaseRepository)
	return srv.Run()
}