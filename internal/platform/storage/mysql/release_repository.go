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
	releaseSQLStruct := sqlbuilder.NewStruct(new(SqlRelease))

	// build the query with the table and the release fields
	query, args := releaseSQLStruct.InsertInto(sqlReleaseTable, SqlRelease{
		ID:          release.ID().String(),
		Title:       release.Title().String(),
		Released:    release.Released().String(),
		ResourceUrl: release.ResourceUrl().String(),
		Uri:         release.Uri().String(),
		Year:        release.Year().String(),
	}).Build()

	// execute the query
	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist release on database: %v", err)
	}

	return nil
}

func (r *ReleaseRepository) All(ctx context.Context) ([]apiExample.Release, error) {
	
	sql, args := sqlbuilder.Select("*").From(sqlReleaseTable).Limit(10).Build()
	
	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return []apiExample.Release{}, err
	}
	
	return FromRows(rows)
}
