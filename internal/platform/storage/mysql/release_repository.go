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
	db        *sql.DB
	dbTimeout time.Duration
}

// NewReleaseRepository initializes a MySQL-based implementation of apiExample.ReleaseRepository.
func NewReleaseRepository(db *sql.DB, dbTimeout time.Duration) *ReleaseRepository {
	return &ReleaseRepository{
		db:        db,
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
func (r *ReleaseRepository) GetAll(ctx context.Context, limit int) ([]apiExample.Release, error) {
	return []apiExample.Release{}, nil
}
