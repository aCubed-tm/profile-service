package main

import (
	"context"
	pb "github.com/acubed-tm/profile-service/protofiles"
)

type server struct{}

func (s *server) CreateProfile(_ context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileReply, error) {
	err := CreateProfileForUuid(req.Uuid, req.FirstName, req.LastName, req.Description)
	if err != nil {
		return nil, err
	} else {
		return &pb.CreateProfileReply{}, nil
	}
}

func (s *server) GetProfile(_ context.Context, req *pb.GetProfileRequest) (*pb.GetProfileReply, error) {
	firstName, lastName, description, err := GetProfileByUuid(req.Uuid)
	if err != nil {
		return nil, err
	} else {
		return &pb.GetProfileReply{
			FirstName:   firstName,
			LastName:    lastName,
			Description: description,
		}, nil
	}
}

func (*server) UpdateProfile(_ context.Context, req *pb.UpdateProfileRequest) (*pb.UpdateProfileReply, error) {
	// note: perhaps check if a field is empty/default?
	err := UpdateProfileForUuid(req.Uuid, req.FirstName, req.LastName, req.Description)
	if err != nil {
		return nil, err
	} else {
		return &pb.UpdateProfileReply{}, nil
	}
}
