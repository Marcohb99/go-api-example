package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/huandu/go-sqlbuilder"
	apiExample "github.com/marcohb99/go-api-example/internal"
)

// ReleaseRepository is a MySQL apiExample.ReleaseRepository implementation.
type ReleaseRepository struct {
	db *sql.DB
}

// NewReleaseRepository initializes a MySQL-based implementation of apiExample.ReleaseRepository.
func NewReleaseRepository(db *sql.DB) *ReleaseRepository {
	return &ReleaseRepository{
		db: db,
	}
}

// Save implements the apiExample.ReleaseRepository interface.
func (r *ReleaseRepository) Save(ctx context.Context, release apiExample.Release) error {
	// build struct based on the DTO class sqlRelease
	releaseSQLStruct := sqlbuilder.NewStruct(new(sqlRelease))

	// build the query with the table and the release fields
	query, args := releaseSQLStruct.InsertInto(sqlReleaseTable, sqlRelease{
		ID:          release.ID(),
		Title:       release.Title(),
		Released:    release.Released(),
		ResourceUrl: release.ResourceUrl(),
		Uri:         release.Uri(),
		Year:        release.Year(),
	}).Build()

	// execute the query
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist release on database: %v", err)
	}

	return nil
}
