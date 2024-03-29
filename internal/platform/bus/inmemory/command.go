package inmemory

import (
	"context"
	"github.com/marcohb99/go-api-example/kit/command"
	"log"
)

// CommandBus is an in-memory implementation of the command.Bus.
type CommandBus struct {
	handlers map[command.Type]command.Handler
}

// NewCommandBus initializes a new instance of CommandBus.
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

// Dispatch implements the command.Bus interface.
func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) (interface{}, error) {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return nil, nil
	}

	resultCh := make(chan interface{})

	go func(ch chan<- interface{}) {
		result, err := handler.Handle(ctx, cmd)
		if err != nil {
			log.Printf("Error while handling %s - %s\n", cmd.Type(), err)
			resultCh <- err
		}
		ch <- result
	}(resultCh)

	response := <-resultCh

	switch response.(type) {
	case error:
		return nil, response.(error)
	default:
		return response, nil
	}
}

// Register implements the command.Bus interface.
func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handlers[cmdType] = handler
}
