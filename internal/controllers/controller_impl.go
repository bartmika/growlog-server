package controllers

import (
	"context"
	// "io"

	"github.com/golang/protobuf/ptypes/empty"
	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/bartmika/growlog-server/proto"
	"github.com/bartmika/growlog-server/internal/models"
)

type ControllerImpl struct {
	tenantRepo models.TenantRepository
	userRepo models.UserRepository
	pb.GrowLogServer
}

func (s *ControllerImpl) Register(ctx context.Context, in *pb.RegistrationInfo) (*pb.User, error) {
	return nil, nil
}

func (s *ControllerImpl) Login(ctx context.Context, in *pb.UserCredential) (*pb.User, error) {
	return nil, nil
}

func (s *ControllerImpl) InsertRow(ctx context.Context, in *pb.TimeSeriesDatum) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *ControllerImpl) InsertRows(stream pb.GrowLog_InsertRowsServer) error {
	return nil
}

func (s *ControllerImpl) Select(in *pb.Filter, stream pb.GrowLog_SelectServer) error {
	return nil
}
