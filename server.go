package main

import (
	"context"
	pb "github.com/acubed-tm/profile-service/protofiles"
)

type server struct{}

func (*server) UpdateProfile(ctx context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileReply, error) {
	// note: perhaps check if a field is empty/default?
	err := UpdateProfileForUuid(ctx, req.Uuid, req.FirstName, req.LastName)
	if err != nil {
		return nil, err
	} else {
		return &pb.UpdateProfileReply{Success: true}, nil
	}
}
