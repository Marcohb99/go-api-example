package mysql

import (
	"context"
	"errors"
	"github.com/marcohb99/go-api-example/internal/platform/factory/factorymocks"
	"github.com/stretchr/testify/mock"
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
	mock.ExpectExec("INSERT INTO releases (uuid, title, released, resource_url, uri, year) VALUES (?, ?, ?, ?, ?, ?)").
		WithArgs(id, title, released, resourceUrl, uri, year).WillReturnError(errors.New("something-failed"))

	// when saving the release
	factory := new(factorymocks.ReleaseFactory)

	repo := NewReleaseRepository(db, 1*time.Millisecond, factory)
	err = repo.Save(context.Background(), release)

	// then the mocked db should be called
	assert.NoError(t, mock.ExpectationsWereMet())
	// and the factory should not be called
	factory.AssertNotCalled(t, "BuildMany")
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
	mock.ExpectExec("INSERT INTO releases (uuid, title, released, resource_url, uri, year) VALUES (?, ?, ?, ?, ?, ?)").
		WithArgs(id, title, released, resourceUrl, uri, year).WillReturnResult(sqlMock.NewResult(0, 1))

	// when saving the release
	factory := new(factorymocks.ReleaseFactory)
	repo := NewReleaseRepository(db, 1*time.Millisecond, factory)
	err = repo.Save(context.Background(), release)

	// then the mocked db should be called
	assert.NoError(t, mock.ExpectationsWereMet())
	// and the factory should not be called
	factory.AssertNotCalled(t, "BuildMany")
	// and no error should be returned
	assert.NoError(t, err)
}

func Test_ReleaseRepository_GetAll_Success(t *testing.T) {
	r1, _ := apiExample.NewRelease(
		"0fc0a540-e84d-4fe7-9868-ee6157562d4b",
		"The Dark Side Of The Moon",
		"1973-03-01",
		"https://api.discogs.com/releases/249504",
		"https://www.discogs.com/Pink-Floyd-The-Dark-Side-Of-The-Moon/release/249504",
		"1973")
	r2, _ := apiExample.NewRelease(
		"0fc0a540-e84d-4fe7-9868-ee6157562d4b",
		"Ultra Mono",
		"2020-01-01",
		"https://api.discogs.com/releases/1809205",
		"https://www.discogs.com/master/1809205-Idles-Ultra-Mono",
		"2020")
	releases := []apiExample.Release{r1, r2}

	// given a mocked db
	db, dbMock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	require.NoError(t, err)
	limit := 10
	dbMock.ExpectQuery("SELECT * FROM releases LIMIT 10").
		WillReturnRows(sqlMock.NewRows([]string{"uuid", "title", "released", "resource_url", "uri", "year"}))

	// and a mocked factory
	factory := new(factorymocks.ReleaseFactory)
	factory.On("BuildMany", mock.Anything).Return(releases, nil)

	// when saving the release
	repo := NewReleaseRepository(db, 1*time.Millisecond, factory)
	result, err := repo.GetAll(context.Background(), limit)

	// then the mocked db should be called
	assert.NoError(t, dbMock.ExpectationsWereMet())

	// and the factory should be called
	factory.AssertCalled(t, "BuildMany", mock.Anything)
	assert.Equal(t, releases, result)

	// and no error should be returned
	assert.NoError(t, err)
}

func Test_ReleaseRepository_GetAll_Error(t *testing.T) {
	// given a mocked db
	db, mock, err := sqlMock.New(sqlMock.QueryMatcherOption(sqlMock.QueryMatcherEqual))
	require.NoError(t, err)
	limit := 10
	mock.ExpectQuery("SELECT * FROM releases LIMIT 10").
		WillReturnError(errors.New("something failed"))

	// when saving the release
	factory := new(factorymocks.ReleaseFactory)
	repo := NewReleaseRepository(db, 1*time.Millisecond, factory)
	_, err = repo.GetAll(context.Background(), limit)

	// then the mocked db should be called
	assert.NoError(t, mock.ExpectationsWereMet())

	// and the factory should not be called
	factory.AssertNotCalled(t, "BuildMany")

	// and no error should be returned
	assert.Error(t, err)
}
