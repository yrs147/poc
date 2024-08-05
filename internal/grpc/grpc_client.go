package grpc

import (
	"context"
	"log"

	"github.com/poc/internal/tls"
	pb "github.com/poc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address = "localhost:50051"
)

func InitClient() {

	// Load TLS credentials
	tlsConfig, err := tls.LoadClientTLS()
	if err != nil {
		log.Fatalf("failed to load ca certs: %v", err)
	}

	// Set up a connection to the server.
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMyServiceClient(conn)

	// Contact the server and print out its response.
	name := "Yash"
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
