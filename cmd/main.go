package main

import (
	"net"

	"google.golang.org/grpc"
	pb "slient.util/generate/proto"

	"slient.util/service"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterSltUtilServiceServer(s, service.NewServer())

	if err = s.Serve(l); err != nil {
		panic(err)
	}
}
