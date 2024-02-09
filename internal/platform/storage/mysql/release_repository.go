package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	apiExample "github.com/marcohb99/go-api-example/internal"
)

// ReleaseRepository is a MySQL apiExample.ReleaseRepository implementation.
type ReleaseRepository struct {
	db *sql.DB
	dbTimeout time.Duration
}

// NewReleaseRepository initializes a MySQL-based implementation of apiExample.ReleaseRepository.
func NewReleaseRepository(db *sql.DB, dbTimeout time.Duration) *ReleaseRepository {
	return &ReleaseRepository{
		db: db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the apiExample.ReleaseRepository interface.
func (r *ReleaseRepository) Save(ctx context.Context, release apiExample.Release) error {
	// build struct based on the DTO class sqlRelease
	releaseSQLStruct := sqlbuilder.NewStruct(new(sqlRelease))

	// build the query with the table and the release fields
	query, args := releaseSQLStruct.InsertInto(sqlReleaseTable, sqlRelease{
		ID:          release.ID().String(),
		Title:       release.Title().String(),
		Released:    release.Released().String(),
		ResourceUrl: release.ResourceUrl().String(),
		Uri:         release.Uri().String(),
		Year:        release.Year().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	// execute the query
	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist release on database: %v", err)
	}

	return nil
}

// All implements the apiExample.ReleaseRepository interface.
func (r *ReleaseRepository) All(ctx context.Context) ([]apiExample.Release, error) {
	// build struct based on the DTO class sqlRelease
	
	query, _ := sqlbuilder.Select("*").From(sqlReleaseTable).Build()
	// execute the query
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return []apiExample.Release{},fmt.Errorf("error trying to retrieve releases on database: %v", err)
	}

	result := []apiExample.Release{}
	// this could be moved to a factory
	for rows.Next() {
		var release apiExample.Release
		err := rows.Scan(&release)
		if err != nil {
			return []apiExample.Release{}, fmt.Errorf("error trying to scan release: %v", err)
		}
		result = append(result, release)
	}

	return result, nil
}
