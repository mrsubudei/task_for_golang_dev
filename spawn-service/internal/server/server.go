// package server creates new grpc server
package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/mrsubudei/task_for_golang_dev/spawn-service/internal/api"
	"github.com/mrsubudei/task_for_golang_dev/spawn-service/internal/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"

	"github.com/mrsubudei/task_for_golang_dev/spawn-service/pkg/logger"
	pb "github.com/mrsubudei/task_for_golang_dev/spawn-service/pkg/proto"
)

type GrpcServer struct {
	l logger.Interface
}

func NewGrpcServer(l logger.Interface) *GrpcServer {
	return &GrpcServer{
		l: l,
	}
}

func (gs *GrpcServer) Start(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcAddr := fmt.Sprintf("%s:%s", cfg.Grpc.Host, cfg.Grpc.Port)

	isReady := &atomic.Value{}
	isReady.Store(false)

	// read ca's cert, verify to client's certificate
	caPem, err := os.ReadFile("cert/ca.cert")
	if err != nil {
		return fmt.Errorf("server - Start - ReadFile: %w", err)
	}

	// create cert pool and append ca's cert
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(caPem) {
		return fmt.Errorf("server - Start - AppendCertsFromPEM: %w", err)
	}

	// read server cert & key
	serverCert, err := tls.LoadX509KeyPair("cert/service.crt", "cert/service.key")
	if err != nil {
		return fmt.Errorf("server - Start - LoadX509KeyPair: %w", err)
	}

	// configuration of the certificate what we want to
	conf := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	//create tls certificate
	tlsCredentials := credentials.NewTLS(conf)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("server - Start - Listen: %w", err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(cfg.Grpc.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(cfg.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(cfg.Grpc.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(cfg.Grpc.Timeout) * time.Minute,
		}),
	)

	pb.RegisterSpawnServer(grpcServer, api.NewSpawnServer(gs.l))

	go func() {
		gs.l.Info("GRPC Server is listening on: %s", grpcAddr)
		if err := grpcServer.Serve(l); err != nil {
			gs.l.Fatal("Failed running gRPC server", err)
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		isReady.Store(true)
		gs.l.Info("The service is ready to accept requests")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		gs.l.Info("signal.Notify: %v", v)
	case done := <-ctx.Done():
		gs.l.Info("ctx.Done: %v", done)
	}

	isReady.Store(false)

	grpcServer.GracefulStop()
	gs.l.Info("grpcServer shut down correctly")

	return nil
}
