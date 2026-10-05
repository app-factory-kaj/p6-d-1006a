package handlers

import (
	"context"
	"testing"

	"greeter/internal/gen"
)

func TestGetGreetingWithName(t *testing.T) {
	s := NewServer()
	resp, err := s.GetGreeting(context.Background(), gen.GetGreetingRequestObject{
		Params: gen.GetGreetingParams{Name: "Ada"},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	greeting, ok := resp.(gen.GetGreeting200JSONResponse)
	if !ok {
		t.Fatalf("unexpected response type: %T", resp)
	}
	if greeting.Name != "Ada" {
		t.Errorf("expected name %q, got %q", "Ada", greeting.Name)
	}
	if greeting.Message != "Hello, Ada!" {
		t.Errorf("unexpected message: %q", greeting.Message)
	}
}

func TestGetGreetingWithoutName(t *testing.T) {
	s := NewServer()
	resp, err := s.GetGreeting(context.Background(), gen.GetGreetingRequestObject{
		Params: gen.GetGreetingParams{},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	greeting, ok := resp.(gen.GetGreeting200JSONResponse)
	if !ok {
		t.Fatalf("unexpected response type: %T", resp)
	}
	if greeting.Name != "" {
		t.Errorf("expected no name, got %q", greeting.Name)
	}
	if greeting.Message != defaultGreeting {
		t.Errorf("expected default greeting, got %q", greeting.Message)
	}
}
