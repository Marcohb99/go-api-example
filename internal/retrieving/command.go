package retrieving

import command "github.com/marcohb99/go-api-example/kit"

const ReleaseCommandType command.Type = "command.retrieving.release"

type ReleaseCommand struct{}

func NewReleaseCommand() ReleaseCommand {
	return ReleaseCommand{}
}

func (c ReleaseCommand) Type() command.Type {
	return ReleaseCommandType
}
