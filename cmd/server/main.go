package main

import (
	"github.com/micro/go-micro/v2/registry/kubernetes"
	"github.com/micro/go-micro/v2/service"
	"github.com/micro/go-micro/v2/service/grpc"
)

// NewService returns a new go-micro service pre-initialised for k8s
func NewService(opts ...service.Option) service.Service {
	// create registry and selector
	r := kubernetes.NewRegistry()

	// set the registry and selector
	options := []service.Option{
		service.Registry(r),
	}

	// append user options
	options = append(options, opts...)

	// return a micro.Service
	return grpc.NewService(options...)
}

func main() {
	service := NewService(
		service.Name("server"),
	)
	service.Init()

	err := service.Run()
	if err != nil {
		panic(err)
	}
}
