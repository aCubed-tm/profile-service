package main

import (
	"context"
	pb "github.com/acubed-tm/profile-service/protofiles"
)

type server struct{}

func (*server) UpdateProfile(_ context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileReply, error) {
	// note: perhaps check if a field is empty/default?
	err := UpdateProfileForUuid(req.Uuid, req.FirstName, req.LastName, req.Description)
	if err != nil {
		return nil, err
	} else {
		return &pb.UpdateProfileReply{Success: true}, nil
	}
}
