package main


import (
	"context"
	"fmt"
	proto "github.com/go-micro-test/api/protoc_gen/v1"
	"github.com/micro/go-micro"
)



func main() {
	service := micro.NewService(
			micro.Name("Greeter.client"),
			)

	fmt。Println("test")

	service.Init()

	greeter := proto.NewGreeterService("Greeter", service.Client())

	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rsp.Greeting)
}
