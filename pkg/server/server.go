package server

import (
	"context"

	pb "github.com/hegemone/io/pkg/user"
)

type UserServiceServer struct {
}

func (s *UserServiceServer) UserIsAdmin(context.Context, *pb.User) (*pb.UserReply, error) {
	return &pb.UserReply{IsAdmin: true}, nil
}
