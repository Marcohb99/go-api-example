package mysql

import (
	"context"
	"errors"
	"testing"
	"time"

	sqlMock "github.com/DATA-DOG/go-sqlmock"
	apiExample "github.com/marcohb99/go-api-example/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ReleaseRepository_Save_RepositoryError(t *testing.T) {
	// given a release
	id := "0fc0a540-e84d-4fe7-9868-ee6157562d4b"
	title := "The Dark Side Of The Moon"
	released := "1973-03-01"
	resourceUrl := "https://api.discogs.com/releases/249504"
	uri := "https://www.discogs.com/Pink-Floyd-The-Dark-Side-Of-The-Moon/release/249504"
	year := "1973"
	release, err := apiExample.NewRelease(id, title, released, resourceUrl, uri, year)
	require.NoError(t, err)

	// and a mocked db
	db, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	require.NoError(t, err)
	mock.ExpectExec("INSERT INTO releases (id, title, released, resource_url, uri, year) VALUES (?, ?, ?, ?, ?, ?)").
		WithArgs(id, title, released, resourceUrl, uri, year).WillReturnError(errors.New("something-failed"))

	// when saving the release
	repo := NewReleaseRepository(db, 1*time.Millisecond)
	err = repo.Save(context.Background(), release)

	// then the mocked db should be called
	assert.NoError(t, mock.ExpectationsWereMet())
	// and an error should be returned
	assert.Error(t, err)
}

func Test_ReleaseRepository_Save_Success(t *testing.T) {
	// given a release
	id := "0fc0a540-e84d-4fe7-9868-ee6157562d4b"
	title := "The Dark Side Of The Moon"
	released := "1973-03-01"
	resourceUrl := "https://api.discogs.com/releases/249504"
	uri := "https://www.discogs.com/Pink-Floyd-The-Dark-Side-Of-The-Moon/release/249504"
	year := "1973"
	release, err := apiExample.NewRelease(id, title, released, resourceUrl, uri, year)
	require.NoError(t, err)

	// and a mocked db
	db, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	require.NoError(t, err)
	mock.ExpectExec("INSERT INTO releases (id, title, released, resource_url, uri, year) VALUES (?, ?, ?, ?, ?, ?)").
		WithArgs(id, title, released, resourceUrl, uri, year).WillReturnResult(sqlMock.NewResult(0, 1))

	// when saving the release
	repo := NewReleaseRepository(db, 1*time.Millisecond)
	err = repo.Save(context.Background(), release)

	// then the mocked db should be called
	assert.NoError(t, mock.ExpectationsWereMet())
	// and no error should be returned
	assert.NoError(t, err)
}

func Test_ReleaseRepository_All_RepositoryError(t *testing.T) {
	// given a mocked db
	db, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	require.NoError(t, err)
	mock.ExpectExec("SELECT * FROM releases").WillReturnError(errors.New("something-failed"))

	// when retrieving the release
	repo := NewReleaseRepository(db, 1*time.Millisecond)
	result, err := repo.All(context.Background())

	// then the mocked db should be called
	assert.NoError(t, mock.ExpectationsWereMet())
	// and an error should be returned
	assert.Error(t, err)
	assert.Empty(t, result)
}

func Test_ReleaseRepository_All_Success(t *testing.T) {
	// given a mocked db
	db, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	require.NoError(t, err)
	mock.ExpectExec("SELECT * FROM releases").WillReturnResult(sqlMock.NewResult(0, 2))

	// when saving the release
	repo := NewReleaseRepository(db, 1*time.Millisecond)
	_, err = repo.All(context.Background())

	// then the mocked db should be called
	assert.NoError(t, mock.ExpectationsWereMet())
	// and no error should be returned
	assert.NoError(t, err)
}
