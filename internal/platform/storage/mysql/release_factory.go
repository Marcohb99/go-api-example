package mysql

import (
	"database/sql"

	apiExample "github.com/marcohb99/go-api-example/internal"
)

// FromRows convert multiple mysql rows to a release collection
func FromRows(rows *sql.Rows) ([]apiExample.Release, error)  {
	var release SqlRelease
	var result []apiExample.Release

    for rows.Next() {
        err := rows.Scan(&release.ID, &release.Title, &release.Released, &release.ResourceUrl, &release.Uri, &release.Year)
        if err != nil {
			return []apiExample.Release{}, err
		}
		releaseObj, err := apiExample.NewRelease(
			release.ID,
			release.Title,
			release.Released,
			release.ResourceUrl,
			release.Uri,
			release.Year,
		)
		if err != nil {
			return []apiExample.Release{}, err
		}
		result = append(result, releaseObj)
    }

    // Check for errors from iterating over rows
    if err := rows.Err(); err != nil {
        return []apiExample.Release{}, err
    }
	return result, nil
}