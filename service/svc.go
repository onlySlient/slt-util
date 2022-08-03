package service

import (
	"context"

	pb "slient.util/generate/proto"
)

type server struct {
	pb.UnimplementedSltUtilServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloResp, error) {
	return &pb.HelloResp{Resp: "hello"}, nil
}
