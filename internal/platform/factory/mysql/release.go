package mysql

import (
	"database/sql"
	apiExample "github.com/marcohb99/go-api-example/internal"
)

type ReleaseFactory struct {
}

func NewReleaseFactory() *ReleaseFactory {
	return &ReleaseFactory{}
}

// BuildMany creates a slice of releases from the given data.
func (f ReleaseFactory) BuildMany(data interface{}) ([]apiExample.Release, error) {
	rows := data.(*sql.Rows)

	// Initialize an empty slice to hold the users
	var releases []apiExample.Release

	// Iterate over the rows
	for rows.Next() {
		// Create a new User struct
		var id string
		var title string
		var released string
		var resourceUrl string
		var uri string
		var year string

		// Scan the values from the current row into the User struct fields
		if err := rows.Scan(&id, &title, &released, &resourceUrl, &uri, &year); err != nil {
			return []apiExample.Release{}, err
		}

		// Append the User struct to the slice
		release, err := apiExample.NewRelease(id, title, released, resourceUrl, uri, year)
		if err != nil {
			return []apiExample.Release{}, err
		}
		releases = append(releases, release)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return releases, nil
}
