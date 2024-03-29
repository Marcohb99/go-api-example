package creating

import (
	"context"
	"errors"
	"github.com/marcohb99/go-api-example/kit/command"
)

const ReleaseCommandType command.Type = "command.creating.release"

// ReleaseCommand is the command dispatched to create a new release.
type ReleaseCommand struct {
	id          string
	title       string
	released    string
	resourceUrl string
	uri         string
	year        string
}

// NewReleaseCommand creates a new ReleaseCommand.
func NewReleaseCommand(id, title, released, resourceUrl, uri, year string) ReleaseCommand {
	return ReleaseCommand{
		id:          id,
		title:       title,
		released:    released,
		resourceUrl: resourceUrl,
		uri:         uri,
		year:        year,
	}
}

func (c ReleaseCommand) Type() command.Type {
	return ReleaseCommandType
}

// ReleaseCommandHandler is the command handler
// responsible for creating releases.
type ReleaseCommandHandler struct {
	service ReleaseService
}

// NewReleaseCommandHandler initializes a new ReleaseCommandHandler.
func NewReleaseCommandHandler(service ReleaseService) ReleaseCommandHandler {
	return ReleaseCommandHandler{
		service: service,
	}
}

// Handle implements the command.Handler interface.
func (h ReleaseCommandHandler) Handle(ctx context.Context, cmd command.Command) (interface{}, error) {
	createReleaseCmd, ok := cmd.(ReleaseCommand)
	if !ok {
		return nil, errors.New("unexpected command")
	}

	return nil, h.service.CreateRelease(
		ctx,
		createReleaseCmd.id,
		createReleaseCmd.title,
		createReleaseCmd.released,
		createReleaseCmd.resourceUrl,
		createReleaseCmd.uri,
		createReleaseCmd.year,
	)
}
