package mock_spawn_api

import (
	"context"

	pb "github.com/mrsubudei/task_for_golang_dev/spawn-service/pkg/proto"
	"google.golang.org/grpc"
)

type MockSpawnApi struct{}

func NewMockSpawnApi() *MockSpawnApi {
	return &MockSpawnApi{}
}

func (ms *MockSpawnApi) Generate(ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.StringResponse, error) {
	return &pb.StringResponse{
		Str: "",
	}, nil
}
