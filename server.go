package main

import (
	"context"

	pb "github.com/acubed-tm/profile-service/protofiles"
	googleUuid "github.com/google/uuid"
)

type server struct{}

func (s *server) CreateProfile(_ context.Context, req *pb.CreateProfileRequest) (*pb.CreateProfileReply, error) {
	profileUuid := googleUuid.New().String()
	err := CreateProfileForUuid(req.Uuid, profileUuid, req.FirstName, req.LastName, req.Description)
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

func (s *server) CreateOrganizationProfile(_ context.Context, req *pb.CreateOrganizationProfileRequest) (*pb.CreateOrganizationProfileReply, error) {
	profileUuid := googleUuid.New().String()
	err := CreateOrganizationProfileForUuid(req.Uuid, profileUuid, req.DisplayName, req.Description)
	if err != nil {
		return nil, err
	} else {
		return &pb.CreateOrganizationProfileReply{}, nil
	}
}

func (s *server) GetOrganizationProfile(_ context.Context, req *pb.GetOrganizationProfileRequest) (*pb.GetOrganizationProfileReply, error) {
	displayName, description, err := GetOrganizationProfileByUuid(req.Uuid)
	if err != nil {
		return nil, err
	} else {
		return &pb.GetOrganizationProfileReply{
			DisplayName: displayName,
			Description: description,
		}, nil
	}
}

func (*server) UpdateOrganizationProfile(_ context.Context, req *pb.UpdateOrganizationProfileRequest) (*pb.UpdateOrganizationProfileReply, error) {
	// note: perhaps check if a field is empty/default?
	err := UpdateOrganizationProfileForUuid(req.Uuid, req.DisplayName, req.Description)
	if err != nil {
		return nil, err
	} else {
		return &pb.UpdateOrganizationProfileReply{}, nil
	}
}
