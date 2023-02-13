package api

import (
	"math/rand"
	"time"

	pb "github.com/mrsubudei/task_for_golang_dev/spawn-service/pkg/proto"
)

type SpawnServer struct {
	pb.UnimplementedSpawnServer
}

func NewSpawnServer() *SpawnServer {
	return &SpawnServer{}
}

func (ss *SpawnServer) Generate(in *pb.Empty) *pb.StringResponse {
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

	return &pb.StringResponse{Str: string(ans)}
}
