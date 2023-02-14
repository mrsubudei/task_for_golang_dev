// package api implements methods defind in proto file
package api

import (
	"context"
	"math/rand"
	"time"

	"github.com/mrsubudei/task_for_golang_dev/spawn-service/pkg/logger"
	pb "github.com/mrsubudei/task_for_golang_dev/spawn-service/pkg/proto"
)

type SpawnServer struct {
	pb.UnimplementedSpawnServer
	l logger.Interface
}

func NewSpawnServer(l logger.Interface) *SpawnServer {
	return &SpawnServer{
		l: l,
	}
}

func (ss *SpawnServer) Generate(ctx context.Context, in *pb.Empty) (*pb.StringResponse, error) {
	rand.Seed(time.Now().UnixNano())
	ans := make([]byte, 0, 12)

	for i := 0; i < 12; i++ {
		tmp := 0
		random := rand.Intn(3)

		if random == 0 {
			tmp = rand.Intn(10) + 48
		} else if random == 1 {
			tmp = rand.Intn(26) + 65
		} else if random == 2 {
			tmp = rand.Intn(26) + 97
		}

		ans = append(ans, byte(tmp))
	}

	return &pb.StringResponse{
		Str: string(ans),
	}, nil
}
