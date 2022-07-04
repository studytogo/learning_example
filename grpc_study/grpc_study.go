package grpc_study

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	pb "pcCobra/grpc_study/hello"
	"time"
)

func StudyGrpc() {
	go func() {
		time.Sleep(3 * time.Second)
		client()
	}()
	Server()
}

type HYSW struct {
	pb.UnimplementedGreeterServer
}

func (s *HYSW) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello HYSW " + in.GetName()}, nil
}

type GTYR struct {
	pb.UnimplementedGreeterServer
}

//func (s *GTYR) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//	return &pb.HelloReply{Message: "Hello GTYR " + in.GetName()}, nil
//}

func client() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "zzg"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func Server1() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println("888888")
	var m HYSW = HYSW{}

	gsvr := grpc.NewServer()
	pb.RegisterGreeterServer(gsvr, &m)

	log.Printf("server listening at %v", listen.Addr())
	if err := gsvr.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
