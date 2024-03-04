package apiExample

import event "github.com/marcohb99/go-api-example/kit/events"

const ReleaseCreatedEventType event.Type = "events.release.created"

type ReleaseCreatedEvent struct {
	event.BaseEvent
	id          string
	title       string
	released    string
	resourceUrl string
	uri         string
	year        string
}

func NewReleaseCreatedEvent(id, title, released, resourceUrl, uri, year string) ReleaseCreatedEvent {
	return ReleaseCreatedEvent{
		id:          id,
		title:       title,
		released:    released,
		resourceUrl: resourceUrl,
		uri:         uri,
		year:        year,

		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e ReleaseCreatedEvent) Type() event.Type {
	return ReleaseCreatedEventType
}

func (e ReleaseCreatedEvent) ReleaseID() string {
	return e.id
}

func (e ReleaseCreatedEvent) ReleaseTitle() string {
	return e.title
}

func (e ReleaseCreatedEvent) ReleaseReleased() string {
	return e.released
}

func (e ReleaseCreatedEvent) ReleaseResourceUrl() string {
	return e.resourceUrl
}

func (e ReleaseCreatedEvent) ReleaseUri() string {
	return e.uri
}

func (e ReleaseCreatedEvent) ReleaseYear() string {
	return e.year
}
