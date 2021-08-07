package internal

import (
	"context"
	// "io"

	"github.com/golang/protobuf/ptypes/empty"
	// tspb "github.com/golang/protobuf/ptypes/timestamp"

	pb "github.com/bartmika/growlog-server/proto"
)

type GrowLogServerImpl struct {
	pb.GrowLogServer
}

func (s *GrowLogServerImpl) InsertRow(ctx context.Context, in *pb.TimeSeriesDatum) (*empty.Empty, error) {
	return &empty.Empty{}, nil
}

func (s *GrowLogServerImpl) InsertRows(stream pb.GrowLog_InsertRowsServer) error {
	return nil
}

func (s *GrowLogServerImpl) Select(in *pb.Filter, stream pb.GrowLog_SelectServer) error {
	return nil
}
