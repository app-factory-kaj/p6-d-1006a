package handlers

import (
	"context"

	"greeter/internal/gen"
)

const defaultGreeting = "Hello, World!"

// Server implements gen.StrictServerInterface.
type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetGreeting(ctx context.Context, request gen.GetGreetingRequestObject) (gen.GetGreetingResponseObject, error) {
	name := request.Params.Name

	if name == "" {
		return gen.GetGreeting200JSONResponse{
			Message: defaultGreeting,
		}, nil
	}

	return gen.GetGreeting200JSONResponse{
		Name:    name,
		Message: "Hello, " + name + "!",
	}, nil
}
