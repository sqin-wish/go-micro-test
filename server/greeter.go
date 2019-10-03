package main

import (
	proto "github.com/go-micro-test/api/protoc_gen/v1"
	"github.com/micro/go-micro"
	"log"
	rpc "github.com/go-micro-test/hanlder"
)

type Greeter struct{}

func main() {
	service := micro.NewService(
			micro.Name("Greeter"),
			micro.Version("latest"),
		)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(rpc.Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}