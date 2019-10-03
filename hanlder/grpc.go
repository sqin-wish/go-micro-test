package hanlder

import (

	"context"
	proto "github.com/go-micro-test/api/protoc_gen/v1"
)

type Greeter struct{}


func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

