package api_spawn

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/mrsubudei/task_for_golang_dev/spawn-service/pkg/proto"
	"github.com/mrsubudei/task_for_golang_dev/users-service/internal/config"
)

func NewClient(cfg *config.Config) (pb.SpawnClient, error) {
	grpcAddr := fmt.Sprintf("%s:%s", cfg.SpawnApi.Host, cfg.SpawnApi.Port)

	caCert, err := ioutil.ReadFile("cert/ca.cert")
	if err != nil {
		return nil, fmt.Errorf("api_spawn - NewClient - ReadFile: %w", err)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(caCert); !ok {
		return nil, fmt.Errorf("api_spawn - NewClient - AppendCertsFromPEM: %w", err)
	}

	//read client cert
	clientCert, err := tls.LoadX509KeyPair("cert/service.pem", "cert/service.key")
	if err != nil {
		return nil, fmt.Errorf("api_spawn - NewClient - LoadX509KeyPair: %w", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	tlsCredential := credentials.NewTLS(config)
	conn, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(tlsCredential))
	if err != nil {
		return nil, fmt.Errorf("api_spawn - NewClient - Dial: %w", err)
	}
	defer conn.Close()
	client := pb.NewSpawnClient(conn)

	return client, nil
}