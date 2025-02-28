package service

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	v1 "server/api/helloworld/v1"
	"server/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	_, span := otel.Tracer("greeter").Start(ctx, "SayHello")
	defer span.End()
	span.SetAttributes(attribute.KeyValue{
		Key:   "name",
		Value: attribute.StringValue(in.Name),
	})
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
