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
	factory   apiExample.ReleaseFactory
}

// NewReleaseRepository initializes a MySQL-based implementation of apiExample.ReleaseRepository.
func NewReleaseRepository(db *sql.DB, dbTimeout time.Duration, factory apiExample.ReleaseFactory) *ReleaseRepository {
	return &ReleaseRepository{
		db:        db,
		dbTimeout: dbTimeout,
		factory:   factory,
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

// GetAll implements the apiExample.ReleaseRepository interface.
func (r *ReleaseRepository) GetAll(ctx context.Context, limit int) ([]apiExample.Release, error) {
	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	// execute the query
	result, err := r.db.QueryContext(ctxTimeout, "SELECT * FROM "+sqlReleaseTable+" LIMIT ?", limit)
	if err != nil {
		return []apiExample.Release{}, fmt.Errorf("error trying to get releases from database: %v", err)
	}

	// FACTORY
	defer result.Close()

	data, err := r.factory.BuildMany(result)
	if err != nil {
		return []apiExample.Release{}, fmt.Errorf("error trying to build releases from database: %v", err)
	}
	return data, nil
}
