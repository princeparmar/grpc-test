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

	// create local listener on port 7000
	netListener := getNetListener(7000)

	// creating grpc server
	gRPCServer := _grpc.NewServer()

	// register our method with grpc server
	s := new(validate.UnimplementedPingServer)
	validate.RegisterPingServer(gRPCServer, s)

	// start client routine
	go startClient()

	// start the server
	if err := gRPCServer.Serve(netListener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

func startClient() {
	// create continueous loop for continueous server call on fixed interval
	for {
		time.Sleep(time.Second)

		// create connection with server
		var conn *_grpc.ClientConn
		conn, err := _grpc.Dial(":7000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		// create client and call SayHello method on server
		c := validate.NewPingClient(conn)
		response, err := c.SayHello(context.Background(), &empty.Empty{})
		if err != nil {
			log.Fatalf("Error when calling SayHello: %s", err)
		}
		log.Printf("Response from server: %s", response.Greeting)
	}
}

// create listener accoring to port
func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
