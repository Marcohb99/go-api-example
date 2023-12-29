package apiExample

// Release is the data structure that represents a release.
type Release struct {
	id          string
	title       string
	released     string
	resourceUrl string
	uri         string
	year        string
}

// NewRelease creates a new release.
func NewRelease(id, title, released, resourceUrl, uri, year string) Release {
	return Release{
		id:          id,
		title:       title,
		released:     released,
		resourceUrl: resourceUrl,
		uri:         uri,
		year:        year,
	}
}

// ID returns the release unique identifier.
func (r Release) ID() string {
	return r.id
}

// Title returns the release title.
func (r Release) Title() string {
	return r.title
}

// Released returns the release date.
func (r Release) Released() string {
	return r.released
}

// Released returns the release resource url on discogs api.
func (r Release) ResourceUrl() string {
	return r.resourceUrl
}

// Released returns the release uri on discogs web.
func (r Release) Uri() string {
	return r.uri
}

// Year returns the release year.
func (r Release) Year() string {
	return r.year
}
