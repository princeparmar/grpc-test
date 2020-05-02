package main

import (
	"context"
	"fmt"
	"grpc/validate/validate"
	"log"
	"net"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	_grpc "google.golang.org/grpc"
)

func main() {
	netListener := getNetListener(7000)
	gRPCServer := _grpc.NewServer()

	s := new(validate.UnimplementedPingServer)
	validate.RegisterPingServer(gRPCServer, s)

	go startClient()
	// start the server
	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func startClient() {
	for {
		time.Sleep(time.Second)
		var conn *_grpc.ClientConn
		conn, err := _grpc.Dial(":7000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()
		c := validate.NewPingClient(conn)
		response, err := c.SayHello(context.Background(), &empty.Empty{})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from server: %s", response.Greeting)
	}
}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
