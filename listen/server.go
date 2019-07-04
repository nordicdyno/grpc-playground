package main

import (
	"context"

	"github.com/nordicdyno/grpc-listen/api"
)

type publisherServer struct {
}

var _ api.PublisherServer = &publisherServer{}

func (s *publisherServer) Topics(ctx context.Context, _ *api.Nope) (*api.TopicNames, error) {
	names := &api.TopicNames{}
	return names, nil
}

func (s *publisherServer) Topic(tn *api.TopicName, stream api.Publisher_TopicServer) error {
	// name := tn.Name

	topics := []string{"FIELD-1: VALUE-1"}
	if err := stream.Send(&api.TopicContent{KV: topics}); err != nil {
		return err
	}
	return nil
}
