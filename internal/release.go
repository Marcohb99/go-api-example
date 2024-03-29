package apiExample

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
	event "github.com/marcohb99/go-api-example/kit/events"
)

// ReleaseRepository defines the expected behaviour for a release storage service.
type ReleaseRepository interface {
	Save(ctx context.Context, release Release) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=ReleaseRepository

// Release ENTITIES
// Release is the data structure that represents a release.
type Release struct {
	id          ReleaseID
	title       ReleaseTitle
	released    ReleaseDate
	resourceUrl ReleaseResourceUrl
	uri         ReleaseURI
	year        ReleaseYear

	events []event.Event
}

// NewRelease creates a new release.
func NewRelease(id, title, released, resourceUrl, uri, year string) (Release, error) {
	idVO, err := NewReleaseID(id)
	if err != nil {
		return Release{}, err
	}

	titleVO, err := NewReleaseTitle(title)
	if err != nil {
		return Release{}, err
	}

	releasedVO, err := NewReleaseDate(released)
	if err != nil {
		return Release{}, err
	}

	resourceUrlVO, err := NewReleaseResourceUrl(resourceUrl)
	if err != nil {
		return Release{}, err
	}

	uriVO, err := NewReleaseURI(uri)
	if err != nil {
		return Release{}, err
	}

	yearVO, err := NewReleaseYear(year)
	if err != nil {
		return Release{}, err
	}

	release := Release{
		id:          idVO,
		title:       titleVO,
		released:    releasedVO,
		resourceUrl: resourceUrlVO,
		uri:         uriVO,
		year:        yearVO,
	}

	release.Record(NewReleaseCreatedEvent(
		idVO.String(),
		titleVO.String(),
		releasedVO.String(),
		resourceUrlVO.String(),
		uriVO.String(),
		yearVO.String(),
	))

	return release, nil
}

// Record records a new domain event.
func (r *Release) Record(evt event.Event) {
	r.events = append(r.events, evt)
}

// PullEvents returns all the recorded domain events.
func (r Release) PullEvents() []event.Event {
	evt := r.events
	r.events = []event.Event{}

	return evt
}

// ID returns the release unique identifier.
func (r Release) ID() ReleaseID {
	return r.id
}

// Title returns the release title.
func (r Release) Title() ReleaseTitle {
	return r.title
}

// Released returns the release date.
func (r Release) Released() ReleaseDate {
	return r.released
}

// Released returns the release resource url on discogs api.
func (r Release) ResourceUrl() ReleaseResourceUrl {
	return r.resourceUrl
}

// Released returns the release uri on discogs web.
func (r Release) Uri() ReleaseURI {
	return r.uri
}

// Year returns the release year.
func (r Release) Year() ReleaseYear {
	return r.year
}

// VALUE OBJECTS

// ID
var ErrInvalidReleaseID = errors.New("invalid Release ID")

// ReleaseID represents the release's unique identifier.
type ReleaseID struct {
	value string
}

func NewReleaseID(value string) (ReleaseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return ReleaseID{}, fmt.Errorf("%w: %s", ErrInvalidReleaseID, value)
	}

	return ReleaseID{
		value: v.String(),
	}, nil
}

func (id ReleaseID) String() string {
	return id.value
}

// TITLE
var ErrEmptyReleaseTitle = errors.New("empty Release title")

// ReleaseID represents the release's title.
type ReleaseTitle struct {
	value string
}

func NewReleaseTitle(value string) (ReleaseTitle, error) {
	if value == "" {
		return ReleaseTitle{}, fmt.Errorf("%w: %s", ErrEmptyReleaseTitle, value)
	}

	return ReleaseTitle{
		value: value,
	}, nil
}

func (title ReleaseTitle) String() string {
	return title.value
}

// RELEASED
var ErrEmptyReleaseDate = errors.New("empty Release date")
var ErrInvalidReleaseDate = errors.New("invalid Release date, it must be a valid date")

// ReleaseID represents the release's title.
type ReleaseDate struct {
	value string
}

func NewReleaseDate(value string) (ReleaseDate, error) {
	if value == "" {
		return ReleaseDate{}, fmt.Errorf("%w: %s", ErrEmptyReleaseDate, value)
	}

	timeLayout := "2006-01-02"
	if _, err := time.Parse(timeLayout, value); err != nil {
		return ReleaseDate{}, fmt.Errorf("%w: %s", ErrInvalidReleaseDate, value)
	}

	return ReleaseDate{
		value: value,
	}, nil
}

func (d ReleaseDate) String() string {
	return d.value
}

// RESOURCE URL
var ErrEmptyReleaseResourceUrl = errors.New("empty Release resource URL")
var ErrInvalidReleaseResourceUrl = errors.New("invalid Release resource URL: it must be a valid url")

// ReleaseID represents the release's title.
type ReleaseResourceUrl struct {
	value string
}

func NewReleaseResourceUrl(value string) (ReleaseResourceUrl, error) {
	if value == "" {
		return ReleaseResourceUrl{}, fmt.Errorf("%w: %s", ErrEmptyReleaseResourceUrl, value)
	}

	if _, err := url.ParseRequestURI(value); err != nil {
		return ReleaseResourceUrl{}, fmt.Errorf("%w: %s", ErrInvalidReleaseResourceUrl, value)
	}

	return ReleaseResourceUrl{
		value: value,
	}, nil
}

func (r ReleaseResourceUrl) String() string {
	return r.value
}

// URI
var ErrEmptyReleaseURI = errors.New("empty Release URI")
var ErrInvalidReleaseURI = errors.New("invalid Release URI: it must be a valid url")

// ReleaseID represents the release's title.
type ReleaseURI struct {
	value string
}

func NewReleaseURI(value string) (ReleaseURI, error) {
	if value == "" {
		return ReleaseURI{}, fmt.Errorf("%w: %s", ErrEmptyReleaseURI, value)
	}

	if _, err := url.ParseRequestURI(value); err != nil {
		return ReleaseURI{}, fmt.Errorf("%w: %s", ErrInvalidReleaseURI, value)
	}

	return ReleaseURI{
		value: value,
	}, nil
}

func (u ReleaseURI) String() string {
	return u.value
}

// Year
var ErrEmptyReleaseYear = errors.New("empty Release year")
var ErrInvalidReleaseYear = errors.New("invalid Release year, it must be a number")

// ReleaseID represents the release's title.
type ReleaseYear struct {
	value string
}

func NewReleaseYear(value string) (ReleaseYear, error) {
	if value == "" {
		return ReleaseYear{}, fmt.Errorf("%w: %s", ErrEmptyReleaseYear, value)
	}

	if _, err := strconv.Atoi(value); err != nil {
		return ReleaseYear{}, fmt.Errorf("%w: %s", ErrInvalidReleaseYear, value)
	}

	return ReleaseYear{
		value: value,
	}, nil
}

func (y ReleaseYear) String() string {
	return y.value
}
