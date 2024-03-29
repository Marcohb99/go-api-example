package retrieving

import (
	"context"
	"errors"
	apiExample "github.com/marcohb99/go-api-example/internal"
	"github.com/marcohb99/go-api-example/kit/command"
)

const ReleaseCommandType command.Type = "command.retrieving.release"

type ReleaseCommand struct {
	limit int
}

func NewReleaseCommand(limit int) ReleaseCommand {
	return ReleaseCommand{limit: limit}
}

func (c ReleaseCommand) Type() command.Type {
	return ReleaseCommandType
}

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
	getReleaseCmd, ok := cmd.(ReleaseCommand)
	if !ok {
		return []apiExample.Release{}, errors.New("unexpected command")
	}

	return h.service.AllReleases(ctx, getReleaseCmd.limit)
}
